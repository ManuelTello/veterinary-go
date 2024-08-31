var form_errors = {}

function addErrorToList(key, value) {
    if (form_errors[key]) {
        form_errors[key] = value
    } else {
        const entries = Object.entries(form_errors)
        entries.push([key, value])
        form_errors = Object.fromEntries(entries)
    }

    for (const [key, value] of Object.entries(form_errors)) {
        if ($(`#error-${key}`).get().length == 0) {
            $(`#group-${key}`).append(`<div class="form-text text-danger" id="error-${key}">${value}</div>`)
        } else {
            $(`#error-${key}`).text(value)
        }
    }
}

function removeErrorFromList(key) {
    delete form_errors[key]
    $(`#error-${key}`).remove()
}

$("#signup-form").on("submit", (event) => {
    event.preventDefault()

    $("input").get().map((e) => {
        if (e.value == "" || !e.value) {
            addErrorToList(e.name, "Field required")
        }
    })
})

$("input[name=email]").on("keyup", () => {
    const value = $("input[name=email]").val()
    if (value.match(new RegExp(/([a-zA-Z0-9](\_|\.)?)*@([a-zA-z]\.?){1,}/g)) == null) {
        addErrorToList("email", "Invalid e-mail address")
    } else {
        removeErrorFromList("email")
    }
})

$("input[name=username]").on("keyup", () => {
    const value = $("input[name=username]").val()
    if (value.length < 4) {
        addErrorToList("username", "Username must be at least 4 characters long")
    } else {
        removeErrorFromList("username")
    }
})

/*
$("input[name=username]").on("change", () => {
    const value = $("input[name=username]").val()
    if (value == "" || !value) {
        addErrorToList("username", "Field required")
    } else if (value.length <= 4) {
        addErrorToList("username", "Username must be at least 4 characters long")
    } else {
        removeErrorFromList("username")
    }
})
    */

/*
let form_values = {
    "email": $("input[name='email']").val(),
    "username": $("input[name='username']").val(),
    "password": $("input[name='password']").val(),
    "first-name": $("input[name='first-name']").val(),
    "last-name": $("input[name='last-name']").val(),
    "phone-number": $("input[name='phone-number']").val(),
    "alt-number": $("input[name='alt-number']").val(),
    "date-created": new Date().toISOString(),
}
    */

/*
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
})*/