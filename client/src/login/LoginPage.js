// @flow strict

import type { Element } from 'react';
import React, { useCallback, useState } from 'react';

import Button from '../component/Button';
import CredentialInput from './CredentialInput';
import Page from '../component/Page';
import PageItemCentered from '../component/PageItemCentered';
import TextTitle from '../component/TextTitle';

import { fetchCreateAccount, fetchLogin } from '../api/login';

type Props = { onSubmit: () => void };

function LoginPage({ onSubmit }: Props): Element<typeof Page> {
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const createAccount = useCallback(
    async (): Promise<void> => {
      const { error } = await fetchCreateAccount(
        password,
        username
      );
      if (!error) {
        onSubmit();
      }
    },
    [password, username],
  );
  const login = useCallback(
    async (): Promise<void> => {
      const { error } = await fetchLogin(
        password,
        username
      );
      if (!error) {
        onSubmit();
      }
    },
    [password, username],
  );
  return (
    <Page>
      <PageItemCentered>
        <TextTitle>trackhours</TextTitle>
        <CredentialInput
          onChange={setUsername}
          placeholder="Username"
          value={username} />
        <CredentialInput
          onChange={setPassword}
          placeholder="Password"
          password
          value={password} />
        <Button onClick={login}
          padding="8px 0"
          width="300px">
          Login
        </Button>
        <Button onClick={createAccount}
          padding="8px 0"
          width="300px">
          Create Account
        </Button>
      </PageItemCentered>
    </Page>
  );
};

export default React.memo<Props>(LoginPage);
