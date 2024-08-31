$("input[name=username]").on("change", () => {
    const username_value = $("input[name=username]").val()
    const body = {
        "username": username_value
    }

    $.ajax("/api/v1/helper/repeatedusername", {
        async: true,
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(body),
        success: (data, status, xhr) => {
            console.log(data)
        },
    })
})

$("input[name=email]").on("change", () => {
    const email_value = $("input[name=email]").val()
    const body = {
        "email": email_value,
    }

    $.ajax("/api/v1/helper/repeatedemail", {
        async: true,
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify(body),
        success: (data, status, xhr) => {
            console.log("email")
        }
    })
})