import {fetchImpl} from './baseAPI';

export const fetchCreateAccount = async (
  password,
  username,
) => (
  await fetchImpl(
    '/api/account_creation',
    {
      password,
      username,
    },
  )
);

export const fetchLogin = async (
  password,
  username,
) => (
  await fetchImpl(
    '/api/login',
    {
      password,
      username,
    },
  )
);
