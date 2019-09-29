// @flow strict

import type {Element, Node} from 'react';
import React from 'react';
import styled from 'styled-components'

const Root = styled.div`
  align-items: center;
  background-color: #202020;
  box-sizing: border-box;
  color: white;
  display: flex;
  flex-direction: column;
  height: fit-content;
  position: relative;
  left: 50%;
  top: 50%;
  transform: perspective(1px) translateY(-50%) translateX(-50%);
`;

type Props = {
  children?: Node,
};

function PageItemCentered(
  {children}: Props,
): Element<typeof Root> {
  return (
    <Root>
      {children}
    </Root>
  );
};

export default React.memo<Props>(PageItemCentered);
