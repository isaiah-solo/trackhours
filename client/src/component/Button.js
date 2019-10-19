// @flow strict

import type {Element, Node} from 'react';
import React from 'react';
import styled from 'styled-components'

type Props = {
  children: Node,
  onClick: () => (void | Promise<void>),
  padding: string,
  width: string,
};

function Button({
  children,
  onClick,
  padding = '0',
  width = '100%'
}: Props): Element<typeof Root> {
  return (
    <Root onClick={onClick}
      padding={padding}
      width={width}>
      {children}
    </Root>
  );
};

const Root = styled.div`
  background-color: #ee0060;
  border-radius: 4px;
  box-sizing: border-box;
  color: white;
  cursor: pointer;
  font-size: 24px;
  font-weight: bold;
  margin-top: 16px;
  padding: ${({padding}) => padding};
  text-align: center;
  user-select: none;
  width: ${({width}) => width};
`;

export default React.memo<Props>(Button);
