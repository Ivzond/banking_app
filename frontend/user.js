let userId;
let isTransactionFormVisible = false
let isTransactionsListVisible = false

document.addEventListener("DOMContentLoaded", function () {
    // Load user information from local storage
    const userInfo = JSON.parse(localStorage.getItem('userInfo'));
    userId = userInfo.ID;
    const username = userInfo.Username;
    const email = userInfo.Email;
    const accounts = userInfo.Accounts; // Array of user's accounts

    // Display user information on the page
    document.getElementById('user-username').innerText = username;
    document.getElementById('user-username-info').innerText = username;
    document.getElementById('user-email').innerText = email;

    // Display user accounts on the page after user information is loaded
    displayUserAccounts(accounts);
});


function displayUserAccounts(accounts) {
    const accountsContainer = document.getElementById('user-accounts');

    // Clear existing content in the container
    accountsContainer.innerHTML = '';

    if (accounts && accounts.length > 0) {
        const accountsList = document.createElement('ul');

        // Loop through each account and create list items
        accounts.forEach(account => {
            const listItem = document.createElement('li');
            listItem.style.marginBottom = '15px';
            listItem.innerHTML = `<strong>Account ID:</strong> ${account.ID}, <strong>Name:</strong> ${account.Name}, <strong>Balance:</strong> ${account.Balance}`;
            accountsList.appendChild(listItem);
        });

        // Append the list of accounts to the container
        accountsContainer.appendChild(accountsList);
    } else {
        // Display a message if the user has no accounts
        accountsContainer.innerHTML = '<p>No accounts found.</p>';
    }
}


function logout() {
    // Clear local storage and redirect to the authorization page
    localStorage.clear();
    window.location.href = 'index.html';
}


function showTransactionForm() {
    // Toggle the visibility of the transaction form
    isTransactionFormVisible = !isTransactionFormVisible
    const transactionForm = document.getElementById('transaction-form');
    transactionForm.style.display = isTransactionFormVisible ? 'block' : 'none';

    hideTransactionList();
}

function hideTransactionList() {
    isTransactionsListVisible = false;
    const transactionsList = document.getElementById('transactions-list');
    transactionsList.style.display = 'none';
}


function makeTransaction() {
    const from = parseInt(document.getElementById('transaction-from').value);
    const to = parseInt(document.getElementById('transaction-to').value);
    const amount = parseInt(document.getElementById('transaction-amount').value);

    fetch('http://localhost:8080/transaction', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('jwt'),
        },
        body: JSON.stringify({ userID: userId, from, to, amount }),
    })
        .then(response => {
            // Display the response message based on the status code
            displayResponseMessage(response.status);

            return response.json()
        })
        .then(data => {
            // Handle the response data, update UI accordingly
            console.log(data);
        })
        .catch(error => console.error('Error:', error));
}

function displayResponseMessage(statusCode) {
    const messageContainer = document.getElementById('response-message');

    messageContainer.innerHTML = '';

    // Create a message based on the status code
    let message;
    switch (statusCode) {
        case 200:
            message = 'Transaction successful!';
            break;
        case 400:
            message = 'Not enough money on balance';
            break;
        case 401:
            message = 'You are not the owner of this account';
            break;
        case 404:
            message = 'Account not found.';
            break;
        default:
            message = 'Unexpected error. Please try again later.';
    }
    const messageElement = document.createElement('p');
    messageElement.textContent = message;
    messageContainer.appendChild(messageElement);
}

function getUserTransactions() {
    // Toggle the visibility of the transactions list
    isTransactionsListVisible = !isTransactionsListVisible;
    const transactionsList = document.getElementById('transactions-list');
    transactionsList.style.display = isTransactionsListVisible ? 'block' : 'none';

    hideTransactionForm();

    fetch(`http://localhost:8080/transaction/${userId}`, {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('jwt'),
        },
    })
        .then(response => response.json())
        .then(data => {
            // Handle the response data, update UI accordingly
            console.log(data);
            displayTransactions(data.data);
        })
        .catch(error => console.error('Error:', error));
}

function hideTransactionForm() {
    isTransactionFormVisible = false;
    const transactionForm = document.getElementById('transaction-form');
    transactionForm.style.display = 'none';
}

function displayTransactions(transactions) {
    const transactionsContainer = document.getElementById('transactions-list');

    // Clear existing content in the container
    transactionsContainer.innerHTML = '';

    if (transactions && transactions.length > 0) {
        const transactionsList = document.createElement('ul');

        // Loop through each transaction and create list items
        transactions.forEach(transaction => {
            const listItem = document.createElement('li');
            listItem.innerHTML = `<strong>From:</strong> ${transaction.From}, <strong>To:</strong> ${transaction.To}, <strong>Amount:</strong> ${transaction.Amount}`;
            transactionsList.appendChild(listItem);
        });

        // Append the list of transactions to the container
        transactionsContainer.appendChild(transactionsList);
    } else {
        // Display a message if there are no transactions
        transactionsContainer.innerHTML = '<p>No transactions found.</p>';
    }
}