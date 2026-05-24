function initQuillEditor(elementId, options) {
    if (typeof Quill === 'undefined') {
        console.error('Quill未加载');
        return null;
    }

    options = options || {};
    var editor = new Quill(elementId, {
        theme: 'snow',
        placeholder: options.placeholder || '请输入内容...',
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline', 'strike'],
                ['blockquote', 'code-block'],
                ['link', 'image', 'video'],

                [{ 'header': 1 }, { 'header': 2 }],
                [{ 'list': 'ordered'}, { 'list': 'bullet' }, { 'list': 'check' }],
                [{ 'script': 'sub'}, { 'script': 'super' }],
                [{ 'indent': '-1'}, { 'indent': '+1' }],
                [{ 'direction': 'rtl' }],

                [{ 'size': ['small', false, 'large', 'huge'] }],
                [{ 'header': [1, 2, 3, 4, 5, 6, false] }],

                [{ 'color': [] }, { 'background': [] }],
                [{ 'font': [] }],
                [{ 'align': [] }],

                ['clean']
            ]
        }
    });

    if (options.initialContent) {
        editor.root.innerHTML = options.initialContent;
    }

    var toolbar = editor.getModule('toolbar');
    toolbar.addHandler('image', function() {
        var range = editor.getSelection();
        var input = document.createElement('input');
        input.type = 'file';
        input.accept = 'image/*';
        input.onchange = function() {
            var file = input.files[0];
            if (file) {
                var reader = new FileReader();
                reader.onload = function(e) {
                    editor.insertEmbed(range.index, 'image', e.target.result);
                };
                reader.readAsDataURL(file);
            }
        };
        input.click();
    });

    return editor;
}
