// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import { FormControl, Grid, InputLabel, MenuItem, Select, TextField } from '@mui/material';
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import dayjs from 'dayjs';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { useContext, useEffect, useState } from 'react';
import { formatDate } from '../../utils/formatUtils';
import { Period } from '../../api/resources';
import { MetricsContext } from '../context/MetricsContext';
import styles from './ResourcesToolbar.module.css';


function ResourcesToolbar() {
  const { api, params } = useContext(MetricsContext);
  const [filter, setFilter] = useState(params.filter);

  const handleFilterChange = (event) => {
    api.setPage(0);
    setFilter(event.target.value);
  };
  const handleDateChange = (event) => {
    api.setDate(formatDate(new Date(event)));
    api.setPage(0);
  };
  const handlePeriodChange = (event) => {
    api.setPeriod(event.target.value);
  };

  useEffect(() => {
    const timer = setTimeout(() => {
      api.setFilter(filter);
    }, 500);
    return () => clearTimeout(timer);
  }, [filter, api]);


  // If we have week selected as the period, disable everything but Sundays
  const handleShouldDisableDate = (date) => {
    if (String(params.period) === '1') {
      return date.day() !== 0;
    }
    return false;
  };

  return (
    <>
      <Grid container className={styles.formContainer} gap={3}>
        <Grid item xs={3}>
          <TextField
            data-testid="textFieldTest"
            fullWidth
            label="Filter"
            variant="standard"
            onChange={handleFilterChange}
            value={filter}
          />
        </Grid>
        <Grid item xs={.7}>
          <FormControl data-testid="formControlTest" fullWidth variant="standard">
            <InputLabel shrink={true}>Period</InputLabel>
            <Select
              value={Number(params.period) as Period}
              label="Period"
              onChange={handlePeriodChange}
            >
              <MenuItem value={0}>Day</MenuItem>
              <MenuItem value={1}>Week</MenuItem>
            </Select>
          </FormControl>
        </Grid>
        <Grid item xs={2}>
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DatePicker
              label="Date"
              disableFuture
              onChange={handleDateChange}
              format="YYYY-MM-DD"
              defaultValue={dayjs(params.date)}
              slotProps={{ textField: { variant: 'standard' } }}
              shouldDisableDate={handleShouldDisableDate}
            />
          </LocalizationProvider>
        </Grid>
      </Grid>
    </>
  );
}

export default ResourcesToolbar;
