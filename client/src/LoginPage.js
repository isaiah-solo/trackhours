import React, {useCallback, useState} from 'react';
import styled from 'styled-components'

import CredentialInput from './component/CredentialInput';
import Page from './component/Page';
import PageItemCentered from './component/PageItemCentered';
import TextTitle from './component/TextTitle';

import {fetchCreateAccount, fetchLogin} from './api/login';

function LoginPage() {
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const createAccount = useCallback(
    async () => {
      const {error} = await fetchCreateAccount(
        password,
        username
      );
      if (!error) {
        window.location.reload();
      }
    },
    [password, username],
  );
  const login = useCallback(
    async () => {
      const {error} = await fetchLogin(
        password,
        username
      );
      if (!error) {
        window.location.reload();
      }
    },
    [password, username],
  );
  return (
    <Page>
      <PageItemCentered>
        <Content>
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
          <Button onClick={login}>
            Login
          </Button>
          <Button onClick={createAccount}>
            Create Account
          </Button>
        </Content>
      </PageItemCentered>
    </Page>
  );
};

const Button = styled.div`
  background-color: #ee0060;
  border-radius: 4px;
  box-sizing: border-box;
  color: white;
  cursor: pointer;
  font-size: 24px;
  font-weight: bold;
  margin-top: 16px;
  padding: 8px 0;
  text-align: center;
  user-select: none;
  width: 300px;
`;

const Content = styled.div`
  align-items: center;
  box-sizing: border-box;
  display: flex;
  flex-direction: column;
`;

export default React.memo(LoginPage);
