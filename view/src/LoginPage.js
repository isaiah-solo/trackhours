import React, {useState} from 'react';
import styled from 'styled-components'

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
  width: 100%;
`;

const Input = styled.input`
  border-radius: 4px;
  border-width: 0px;
  box-sizing: border-box;
  font-size: 24px;
  margin-top: 16px;
  padding: 8px;
  width: 300px;
  :focus {
    outline-width: 0;
  }
`;

const Page = styled.div`
  align-items: center;
  background-color: #202020;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  height: 100vh;
  width: 100vw;
`;

const Title = styled.div`
  box-sizing: border-box;
  color: white;
  font-size: 64px;
  font-weight: bold;
  text-align: center;
  width: 100%;
`;

const CredentialInput = (props) => (
  <Input
    {...props}
    autocapitalize="off"
    autocomplete="off"
    autocorrect="off"
    spellcheck="false" />
);

const LoginPage = (props) => {
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const handleSetPassword = (e) =>
    setPassword(e.target.value);
  const handleSetUsername = (e) =>
    setUsername(e.target.value);
  return (
    <Page>
      <Content>
        <Title>
          {'trackhours'}
        </Title>
        <CredentialInput
          onChange={handleSetUsername}
          placeholder="username"
          type="text"
          value={username} />
        <CredentialInput
          onChange={handleSetPassword}
          placeholder="password"
          type="password"
          value={password} />
        <Button
          onClick={() => (
            fetch(
              'http://localhost:8081/api/login',
              {
                body: JSON.stringify({
                  password,
                  username,
                }),
                credentials: 'same-origin',
                method: 'POST',
                mode: 'no-cors',
              },
            )
              .then((response) => {console.log(response.body)})
          )}
        >
          {'Login'}
        </Button>
        <Button
          onClick={() => (
            fetch(
              'http://localhost:8081/api/account_creation',
              {
                body: JSON.stringify({
                  password,
                  username,
                }),
                credentials: 'same-origin',
                method: 'POST',
                mode: 'no-cors',
              },
            )
              .then((response) => {console.log(response.body)})
          )}
        >
          {'Create Account'}
        </Button>
      </Content>
    </Page>
  );
};

export default LoginPage;
