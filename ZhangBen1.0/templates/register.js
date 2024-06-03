window.onload = function () {
   Captcha();
}
function Captcha() {
    fetch("/LGRG/Captcha", {
        method :"GET" ,
        headers : {
            "Content-Type": "application/json",
        },

    })
        .then(response => response.json())
        .then(data => {
            document.getElementById("captcha-id").value = data.captchaid;
            document.getElementById("img_captcha").querySelector("img").src = data.imagedata;
            //   document.querySelector(".captcha-image img").src = "data:image/jpeg;base64," + data.imagedata;

        })
        .catch(error => console.error("Error",error));
}

document.addEventListener("DOMContentLoaded",function () {
    var captchaImg = document.getElementById("xiaxinimg");
    captchaImg.addEventListener("click",function () {
       Captcha();
    });
})

document.getElementById('registrationForm').addEventListener('submit',function (event){

    event.preventDefault();
    var username = document.getElementById('username').value;
    var nickname = document.getElementById('nickname').value;
    var email = document.getElementById('email').value;
    var phone = document.getElementById('phone').value;
    var password = document.getElementById('password').value;
    var confirmpassword = document.getElementById('cfpwd').value;
    var captchaid = document.getElementById('captcha-id').value;
    var captchasolution = document.getElementById('captcha-value').value;
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/LGRG/reg', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            console.log(xhr.responseText);
            window.location.href = '/login';
            alert(xhr.responseText)
        } else if (xhr.readyState === 4) {
            console.error('Error:', xhr.status, xhr.statusText);
            alert(xhr.responseText)
            Captcha();
        }
    };
    var data = JSON.stringify({
        username:username,
        nickname:nickname,
        email:email,
        phone:phone,
        password:password,
        confirmpassword:confirmpassword,
        captchaid:captchaid,
        captchaanswer:captchasolution,
    });

    xhr.send(data);
});

