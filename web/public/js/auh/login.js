(function () {
    $('#loginForm').submit(function (e) {
        e.preventDefault();
        let data = {
            Email: $('#email').val(),
            Password: $('#password').val(),
        }
        $.ajax({
            url: '/auth/login',
            method: 'post',
            data: data,
            success(_d) {
                console.log(_d)
            },
            error(err) {
                console.log(err)
            }
        })
    })
})()