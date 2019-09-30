import React, {useCallback} from 'react';
import styled from 'styled-components'

function CredentialInput({
  onChange,
  password = false,
  placeholder,
  value
}) {
  const onInputChange = useCallback(
    (e) => (
      onChange(e.target.value)
    ),
    [onChange],
  );
  return (
    <Root
      onChange={onInputChange}
      placeholder={placeholder}
      type={password ? 'password' : 'text'}
      value={value} />
  );
};

const Root = styled.input`
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

export default React.memo(CredentialInput);
