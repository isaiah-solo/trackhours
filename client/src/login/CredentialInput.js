// @flow strict

import type {Element} from 'react';
import React, {useCallback} from 'react';
import styled from 'styled-components'

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

type Props = {
  onChange: string => void,
  password?: boolean,
  placeholder: string,
  value: string,
};

function CredentialInput({
  onChange,
  password = false,
  placeholder,
  value
}: Props): Element<typeof Root> {
  const onInputChange = useCallback(
    (e: SyntheticInputEvent<>): void => (
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

export default React.memo<Props>(CredentialInput);
