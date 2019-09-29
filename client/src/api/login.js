// @flow strict

import {fetchImpl} from './baseAPI';

export const fetchCreateAccount = async (
  password: string,
  username: string,
): {} => (
  await fetchImpl(
    '/api/account_creation',
    {
      password,
      username,
    },
  )
);

export const fetchLogin = async (
  password: string,
  username: string,
): {} => (
  await fetchImpl(
    '/api/login',
    {
      password,
      username,
    },
  )
);
