// 通用图片上传工具
// 选择图片 → 自动上传 → 返回URL，各处统一使用

function previewImage(file, imgElement) {
    const reader = new FileReader();
    reader.onload = function(e) {
        imgElement.src = e.target.result;
        imgElement.style.display = 'block';
    };
    reader.readAsDataURL(file);
}

async function uploadImage(file, type) {
    const formData = new FormData();
    formData.append('img', file);
    formData.append('type', type);
    return await axios.post('/api/img/upload', formData);
}

// 配置式自动上传绑定
// config: { inputId, previewId, statusId (可选), type, onSuccess(url), onError() }
function setupImageUpload(config) {
    const input = document.getElementById(config.inputId);
    const preview = document.getElementById(config.previewId);
    const status = config.statusId ? document.getElementById(config.statusId) : null;

    input.addEventListener('change', async function() {
        const file = this.files[0];
        if (!file) return;

        previewImage(file, preview);

        if (status) {
            status.textContent = '上传中...';
            status.style.color = '#999';
        }

        try {
            const response = await uploadImage(file, config.type);
            if (response.data.code === 0 && response.data.data && response.data.data.url) {
                const url = response.data.data.url;
                preview.src = url;
                if (status) {
                    status.textContent = '上传成功';
                    status.style.color = '#67c23a';
                }
                if (config.onSuccess) config.onSuccess(url);
            } else {
                if (status) {
                    status.textContent = '上传失败';
                    status.style.color = '#f56c6c';
                }
                if (config.onError) config.onError(response.data.msg || '上传失败');
            }
        } catch (error) {
            console.error('Upload error:', error);
            if (status) {
                status.textContent = '上传失败';
                status.style.color = '#f56c6c';
            }
            if (config.onError) config.onError('网络错误');
        }
    });
}
