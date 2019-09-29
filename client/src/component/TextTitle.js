// @flow strict

import type {Element, Node} from 'react';
import React from 'react';
import styled from 'styled-components'

const Root = styled.div`
  box-sizing: border-box;
  color: ${({color}) => color};
  font-size: 64px;
  font-weight: bold;
  text-align: center;
  width: 100%;
`;

type Props = {
  children?: Node,
  color?: string,
};

function TextTitle(
  {children, color = 'white'}: Props,
): Element<typeof Root> {
  return (
    <Root color={color}>
      {children}
    </Root>
  );
};

export default React.memo<Props>(TextTitle);
