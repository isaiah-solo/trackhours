const isDev = process.env.NODE_ENV === 'development';
const API_URI = isDev
  ? 'http://localhost:8081'
  : 'http://trackhours.co';

export const fetchCreateAccount = (password, username) => (
  fetch(
    `${API_URI}/api/account_creation`,
    {
      body: JSON.stringify({
        password,
        username,
      }),
      credentials: 'same-origin',
      method: 'POST',
      mode: 'no-cors',
    },
  )
);

export const fetchLogin = (password, username) => (
  fetch(
    `${API_URI}/api/login`,
    {
      body: JSON.stringify({
        password,
        username,
      }),
      credentials: 'same-origin',
      method: 'POST',
      mode: 'no-cors',
    }
  )
);
