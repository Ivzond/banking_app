async function login() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ Username: username, Password: password }),
    });

    const data = await response.json();
    console.log(data);
}

async function register() {
    const regUsername = document.getElementById('regUsername').value;
    const email = document.getElementById('email').value;
    const regPassword = document.getElementById('regPassword').value;

    const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ Username: regUsername, Email: email, Password: regPassword }),
    });

    const data = await response.json();
    console.log(data);
}

