/**
 * PayQRCode - 扫码支付模块
 * 封装创建订单 -> 展示二维码 -> 轮询支付状态 -> 处理结果的完整流程
 *
 * 使用方式(新建订单):
 *   PayQRCode.open({
 *       createOrder: { url: '/api/orders/add', params: { cid: 1, level: '0', price: 10 } },
 *       poll: {
 *           url: '/api/orders/get',
 *           getParams: function(orderId) { return { orders_id: orderId }; },
 *           isPaid: function(data) { return data.orders && data.orders.orderStatus === 1; }
 *       },
 *       price: 10,
 *       modalTitle: '扫码支付',
 *       onSuccess: function(orderId) { window.location.href = '/page/circle/index?cid=' + cid; },
 *       onCancel: function() {}
 *   });
 *
 * 使用方式(已有订单):
 *   PayQRCode.open({
 *       existingOrder: { orderId: 123, getQrcode: '/api/orders/qrcode', qrcodeParam: 'orders_id' },
 *       poll: { ... },
 *       price: 10,
 *       modalTitle: '扫码支付',
 *       unpaidRedirectUrl: '/page/orders/index',
 *       onSuccess: function(orderId) { ... },
 *       onCancel: function() {}
 *   });
 */
var PayQRCode = (function() {
    var pollingTimer = null;
    var currentOrderId = null;
    var currentOptions = null;
    var modalId = 'payQrModal';

    function getModal() {
        return document.getElementById(modalId);
    }

    function ensureModal() {
        if (getModal()) return;

        var html = '' +
            '<div class="modal" id="' + modalId + '">' +
            '  <div class="modal-content">' +
            '    <h3 id="' + modalId + 'Title">扫码支付</h3>' +
            '    <p style="color:#666;font-size:13px;margin-bottom:12px;" id="' + modalId + 'Price"></p>' +
            '    <img id="' + modalId + 'Img" src="" alt="支付二维码" style="display:block;margin:0 auto 12px auto;max-width:200px;border-radius:8px;">' +
            '    <p style="color:#909399;font-size:12px;margin-bottom:12px;" id="' + modalId + 'Hint">请使用手机扫描二维码完成支付</p>' +
            '    <p style="color:#e6a23c;font-size:13px;margin-bottom:12px;display:none;" id="' + modalId + 'Status">等待支付中...</p>' +
            '    <div class="modal-btns">' +
            '      <button class="btn btn-primary" id="' + modalId + 'SuccessBtn">已完成支付</button>' +
            '      <button class="btn btn-secondary" id="' + modalId + 'CancelBtn">取消</button>' +
            '    </div>' +
            '  </div>' +
            '</div>';

        var div = document.createElement('div');
        div.innerHTML = html;
        document.body.appendChild(div.firstElementChild);

        // Bind events
        document.getElementById(modalId + 'SuccessBtn').addEventListener('click', function() {
            stopPolling();
            closeModal();
            if (currentOptions && currentOptions.onSuccess) {
                currentOptions.onSuccess(currentOrderId);
            }
        });

        document.getElementById(modalId + 'CancelBtn').addEventListener('click', function() {
            stopPolling();
            closeModal();
            if (currentOptions && currentOptions.onCancel) {
                currentOptions.onCancel();
            }
        });

        // Click outside modal to close
        getModal().addEventListener('click', function(e) {
            if (e.target === getModal()) {
                stopPolling();
                closeModal();
                if (currentOptions && currentOptions.onCancel) {
                    currentOptions.onCancel();
                }
            }
        });
    }

    function stopPolling() {
        if (pollingTimer) {
            clearInterval(pollingTimer);
            pollingTimer = null;
        }
    }

    function closeModal() {
        stopPolling();
        var modal = getModal();
        if (modal) {
            modal.classList.remove('show');
        }
    }

    function showModal() {
        ensureModal();
        getModal().classList.add('show');
    }

    function startPolling(orderId) {
        var statusEl = document.getElementById(modalId + 'Status');
        statusEl.style.display = 'block';
        statusEl.style.color = '#e6a23c';
        statusEl.textContent = '等待支付中...';

        pollingTimer = setInterval(async function() {
            try {
                var opts = currentOptions;
                var response = await axios.get(opts.poll.url, {
                    params: opts.poll.getParams(orderId)
                });

                if (response.data.code === 0 && response.data.data) {
                    if (opts.poll.isPaid(response.data.data)) {
                        clearInterval(pollingTimer);
                        pollingTimer = null;
                        onPaymentSuccess();
                    }
                }
            } catch (error) {
                // 轮询出错不中断
            }
        }, 3000);
    }

    function onPaymentSuccess() {
        document.getElementById(modalId + 'Title').textContent = '付款成功';
        document.getElementById(modalId + 'Hint').style.display = 'none';
        var statusEl = document.getElementById(modalId + 'Status');
        statusEl.style.display = 'block';
        statusEl.style.color = '#67c23a';
        statusEl.textContent = '付款成功！';
        document.getElementById(modalId + 'Img').style.display = 'none';

        setTimeout(function() {
            closeModal();
            if (currentOptions && currentOptions.onSuccess) {
                currentOptions.onSuccess(currentOrderId);
            }
        }, 1500);
    }

    /**
     * 打开支付弹窗
     * @param {Object} options
     *   options.createOrder: { url, params } - 创建订单的API
     *   options.poll: { url, getParams(orderId), isPaid(data) } - 轮询配置
     *   options.price - 支付金额
     *   options.modalTitle - 弹窗标题（可选，默认"扫码支付"）
     *   options.onSuccess(orderId) - 支付成功回调
     *   options.onCancel() - 取消回调
     *   options.onError(msg) - 错误回调（可选）
     */
    function open(options) {
        currentOptions = options;

        ensureModal();

        // 已有订单模式：直接获取二维码
        if (options.existingOrder) {
            currentOrderId = options.existingOrder.orderId;
            var qrcodeParam = options.existingOrder.qrcodeParam || 'orders_id';
            var qrcodeParams = {};
            qrcodeParams[qrcodeParam] = currentOrderId;

            axios.get(options.existingOrder.getQrcode, {
                params: qrcodeParams
            }).then(function(response) {
                if (response.data.code === 0) {
                    var data = response.data.data;
                    showQrCodeAndPoll(data.qr_code_url);
                } else {
                    if (options.onError) {
                        options.onError(response.data.msg || '获取二维码失败');
                    }
                }
            }).catch(function(error) {
                if (options.onError) {
                    options.onError('网络错误，请稍后重试');
                }
            });
            return;
        }

        // 新建订单模式
        var formData = new FormData();
        var params = options.createOrder.params;
        for (var key in params) {
            if (params.hasOwnProperty(key)) {
                formData.append(key, params[key]);
            }
        }

        axios.post(options.createOrder.url, formData).then(function(response) {
            if (response.data.code === 0) {
                var data = response.data.data;
                currentOrderId = data.orders_id || data.purchase_id;
                showQrCodeAndPoll(data.qr_code_url);
            } else if (response.data.code === 9013) {
                alert('您已有待支付订单，请前往订单页完成支付');
                window.location.href = options.unpaidRedirectUrl || '/page/orders/index';
            } else {
                if (options.onError) {
                    options.onError(response.data.msg || '创建订单失败');
                }
            }
        }).catch(function(error) {
            if (options.onError) {
                options.onError('网络错误，请稍后重试');
            }
        });
    }

    function showQrCodeAndPoll(qrCodeUrl) {
        document.getElementById(modalId + 'Price').textContent = '支付金额：' + currentOptions.price + ' 元';
        document.getElementById(modalId + 'Img').src = qrCodeUrl;
        document.getElementById(modalId + 'Img').style.display = 'block';
        document.getElementById(modalId + 'Title').textContent = currentOptions.modalTitle || '扫码支付';
        document.getElementById(modalId + 'Hint').style.display = 'block';
        document.getElementById(modalId + 'Status').style.display = 'none';
        showModal();

        startPolling(currentOrderId);
    }

    return {
        open: open,
        close: closeModal
    };
})();
