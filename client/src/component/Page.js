// @flow strict

import type {Element, Node} from 'react';
import React from 'react';
import styled from 'styled-components'

type Props = {
  children?: Node,
};

function Page(
  {children}: Props
): Element<typeof Root> {
  return (
    <Root>
      {children}
    </Root>
  );
};

const Root = styled.div`
  background-color: #202020;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: row;
  height: 100vh;
  width: 100vw;
`;

export default React.memo<Props>(Page);
