// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import Tooltip from '@mui/material/Tooltip';
import IconButton from '@mui/material/IconButton';
import HelpOutline from '@mui/icons-material/HelpOutline';

interface Props {
    text: string;
}

const HelpTooltip= ({ text }: Props) => {
  return (
    <Tooltip arrow title={text}>
      <IconButton aria-label='What is this?'>
        <HelpOutline></HelpOutline>
      </IconButton>
    </Tooltip>
  );
};

export default HelpTooltip;
