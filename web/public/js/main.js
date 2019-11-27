var Helper = {
    parseError: function (err) {
        let txt = JSON.parse(err.responseText)
        if (txt) {
            if (txt.msg) {
                return txt.msg
            }
            if (txt.error) {
                return txt.error
            }
            return "Lỗi không xác định. Vui lòng liên hệ admin"
        }
    }
}