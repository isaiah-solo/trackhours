const isDev = process.env.NODE_ENV === 'development';
const API_URI = isDev
  ? 'http://localhost:8081'
  : 'http://trackhours.co';

const fetchImpl = (path, vars) => (
  fetch(
    API_URI + path,
    {
      body: vars != null ? JSON.stringify(vars) : undefined,
      credentials: 'same-origin',
      method: vars != null ? 'POST' : 'GET',
      mode: 'no-cors',
    },
  )
);

export const fetchCreateAccount = (password, username) => (
  fetchImpl(
    '/api/account_creation',
    {
      password,
      username,
    },
  )
);

export const fetchLogin = (password, username) => (
  fetchImpl(
    '/api/login',
    {
      password,
      username,
    },
  )
);
