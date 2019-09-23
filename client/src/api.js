const API_URI = process.env.REACT_ENVIRONMENT === 'dev'
  ? 'http://localhost:8081'
  : 'http://trackhours.co';

export const fetchLogin = (password, username) => {
  return fetch(
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
  );
}
