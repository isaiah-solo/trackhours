// @flow strict

import {fetchImpl} from './baseAPI';

type Response = {
  error: Error,
};

export const fetchCreateAccount = async (
  password: string,
  username: string,
): Promise<Response> => (
  await fetchImpl<Response>(
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
): Promise<Response> => (
  await fetchImpl<Response>(
    '/api/login',
    {
      password,
      username,
    },
  )
);
