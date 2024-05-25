

window.onload = function () {
    var defaultoption = "1";
    var selectElement = document.getElementById("account-book");
    selectElement.value = defaultoption;

    fetchYearData();
    fetchMonthData();
    fetchHistoryData();
    updateCookie();
}

function fetchYearData() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/YearData', true); // Assuming your backend endpoint is /data
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var data = JSON.parse(xhr.responseText);
            var tableBody = document.querySelector('#YearData tbody');
            tableBody.innerHTML = ''; // Clear existing data
            data.forEach(function (row) {
                var tr = document.createElement('tr');
                tr.innerHTML =
                    '<td>' + row.years + '</td>' +
                    '<td>' + row.yearincome + '</td>' +
                    '<td>' + row.yearexpense + '</td>' +
                    '<td>' + row.credit + '</td>'
                ;
                tableBody.appendChild(tr);
            });
        }
    };
    xhr.send();
}

function  fetchMonthData() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/MonthData', true); // Assuming your backend endpoint is /data
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var data = JSON.parse(xhr.responseText);
            var tableBody = document.querySelector('#MonthData tbody');
            tableBody.innerHTML = ''; // Clear existing data
            data.forEach(function (row) {
                var tr = document.createElement('tr');
                tr.innerHTML =
                    //            '<td class="dataid" style="display:none;">' + row.dataid + '</td>' +
                    '<td>' + row.yearmonth + '</td>' +
                    '<td>' + row.monthincome + '</td>' +
                    '<td>' + row.monthexpense + '</td>' +
                    '<td>' + row.credit + '</td>'
                   ;
                tableBody.appendChild(tr);
            });
        }
    };
    xhr.send();
}

function fetchHistoryData() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/HistoryData', true); // Assuming your backend endpoint is /data
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var data = JSON.parse(xhr.responseText);
            var tableBody = document.querySelector('#HistoryData tbody');
            tableBody.innerHTML = ''; // Clear existing data
            data.forEach(function (row) {
                var tr = document.createElement('tr');
                tr.innerHTML =
        //            '<td class="dataid" style="display:none;">' + row.dataid + '</td>' +
                    '<td>' + row.yearmonthday + '</td>' +
                    '<td>' + row.typ + '</td>' +
                    '<td>' + row.money + '</td>' +
                    '<td>' + row.notes + '</td>' +
                    '<td><button onclick="deleteRow(' + row.dataid + ')">Delete</button></td>'
                    ;
                tableBody.appendChild(tr);
            });
        }
    };
    xhr.send();
}
function deleteRow(dataid) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/delete-data', true); // Assuming your delete endpoint is /delete
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            // If deletion is successful, re-fetch data to update the table
            fetchHistoryData();
            fetchMonthData();
            fetchYearData();
        }
    };
    xhr.send(JSON.stringify({ dataid: dataid }));
}

function  updateCookie() {
    var selecteOption = document.getElementById("account-book").value;
    var xhr = new XMLHttpRequest();
    xhr.open("POST","/ZB-set-cookie",true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
    xhr.send(JSON.stringify({ option: selecteOption }));
}


document.getElementById('AddZD').addEventListener('submit',function (event){
    event.preventDefault();
    var book = document.getElementById('account-book').value;
    var typ = document.getElementById('typ').value;
    var inex = document.getElementById('inex').value;
    var money = document.getElementById('money').value;
    var note = document.getElementById('note').value;

    var zhangdan = {
        book:book,
        typ : typ,
        inex : inex,
        money : money,
        note : note,
    }
    var T = {
        method:'POST',
        headers: {
            'Content-Type': 'application/json'
        }, // 指定数据格式为JSON
        body:JSON.stringify(zhangdan)
    }
    fetch('/index/AddZhangDan',T)
        .then(respose => {
            if (!respose.ok) {
                alert(respose)
                throw new Error ('Network response was not ok');
            }
            return respose.json();
        })
        .then(zhangdan => {
            console.log(zhangdan);
        })
        .catch(error => {
            console.error('There was a problem with your fetch operation:', error);
        });
    location.reload();
})
