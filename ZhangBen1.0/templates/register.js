document.getElementById('registrationForm').addEventListener('submit',function (event){

    event.preventDefault();
    var username = document.getElementById('username').value;
    var nickname = document.getElementById('nickname').value;
    var email = document.getElementById('email').value;
    var phone = document.getElementById('phone').value;
    var password = document.getElementById('password').value;
    var confirmpassword = document.getElementById('cfpwd').value
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/LGRG/reg', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log(xhr.responseText);
            alert(xhr.responseText)
        } else if (xhr.readyState === 4) {
            console.error('Error:', xhr.status, xhr.statusText);
            alert(xhr.responseText)
        }
    };
    var data = JSON.stringify({
        username:username,
        nickname:nickname,
        email:email,
        phone:phone,
        password:password,
        confirmpassword:confirmpassword
    });

    xhr.send(data);
});