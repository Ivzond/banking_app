let isLoginFormVisible = false;
let isRegisterFormVisible = false;

function showLoginForm() {
    isLoginFormVisible = !isLoginFormVisible;
    const loginForm = document.getElementById('login-form');
    loginForm.style.display = isLoginFormVisible ? 'block' : 'none';

    // Hide the register form when showing the login form
    hideRegisterForm();
}

function hideLoginForm() {
    isLoginFormVisible = false;
    const loginForm = document.getElementById('login-form');
    loginForm.style.display = 'none';
}

function showRegisterForm() {
    isRegisterFormVisible = !isRegisterFormVisible;
    const registerForm = document.getElementById('register-form');
    registerForm.style.display = isRegisterFormVisible ? 'block' : 'none';

    // Hide the login form when showing the register form
    hideLoginForm();
}

function hideRegisterForm() {
    isRegisterFormVisible = false;
    const registerForm = document.getElementById('register-form');
    registerForm.style.display = 'none';
}

function login() {
    console.log('Login function called')
    const username = document.getElementById('login-username').value;
    const password = document.getElementById('login-password').value;

    fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    })
        .then(response => {
            console.log('Response status:', response.status);

            // Check if the response status is 200 (OK)
            if (response.status === 200) {
                return response.json();
            } else {
                // Display error message to the user
                throw new Error('Wrong username or password');
            }
        })
        .then(data => {
            // Handle the response data, update UI accordingly
            console.log('Response data:', data);

            // Check if the login was successful
            if (data.message === "OK") {
                // Store JWT token and user info in local storage
                localStorage.setItem('jwt', data.jwt);
                localStorage.setItem('userInfo', JSON.stringify(data.data));

                // Redirect to user page
                window.location.href = 'user.html';
            }
        })
        .catch(error => {
            // Display error message to the user
            console.error('Error:', error.message);
            // Update the UI to show the error message to the user
            document.getElementById('login-error').innerText = error.message;
        });
}

function register() {
    const name = document.getElementById('register-name').value;
    const username = document.getElementById('register-username').value;
    const email = document.getElementById('register-email').value;
    const password = document.getElementById('register-password').value;

    fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, username, email, password }),
    })
        .then(response => {
            // Check if the response status is 200 (OK)
            if (response.status === 200) {
                return response.json();
            } else {
                // Display error message to the user
                throw new Error('Registration failed. Please check your input and try again.');
            }
        })
        .then(data => {
            // Handle the response data, update UI accordingly
            console.log(data);

            // Check if the registration was successful
            if (data.message === "OK") {
                // Store JWT token and user info in local storage
                localStorage.setItem('jwt', data.jwt);
                localStorage.setItem('userInfo', JSON.stringify(data.data));
            }
        })
        .catch(error => {
            // Display error message to the user
            console.error('Error:', error.message);
            // Update the UI to show the error message to the user
            document.getElementById('register-error').innerText = error.message;
        });
}
