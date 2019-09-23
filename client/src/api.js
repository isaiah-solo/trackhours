const API_URI = process.env.NODE_ENV === 'development'
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
