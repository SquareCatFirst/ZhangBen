document.getElementById('loginForm').addEventListener('submit',function (event){
    event.preventDefault();
    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/loginTry', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log(xhr.responseText);
            //         alert("登录成功");
            //        xhr.responseText='';
            //      xhr.responseType='document'

            //    var text = xhr.responseText;
            //   text.
            //    document.write(xhr.responseText);
            window.location.href = '/test/index.html';

        } else if (xhr.readyState === 4) {
            console.error('Error:', xhr.status, xhr.statusText);
            alert(xhr.responseText)
        }
    };
    var data = JSON.stringify({ username: username, password: password });
    xhr.send(data);

});