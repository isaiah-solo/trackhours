// @flow strict

export const isDev = process.env.NODE_ENV === 'development';

export const API_URI = isDev
  ? 'http://localhost:8081'
  : 'http://trackhours.co';

export const fetchImpl = async (path: string, vars?: {}): {} => {
  const response = await fetch(
    API_URI + path,
    {
      body: vars != null ? JSON.stringify(vars) : undefined,
      credentials: isDev ? 'include' : undefined,
      method: vars != null ? 'POST' : 'GET',
      mode: 'cors',
    },
  );
  return await response.json();
};
