import {fetchImpl} from './baseAPI';

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
