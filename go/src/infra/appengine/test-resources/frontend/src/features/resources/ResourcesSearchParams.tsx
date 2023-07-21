// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import { useCallback, useContext, useEffect } from 'react';
import { useSearchParams } from 'react-router-dom';
import { MetricsContext } from '../context/MetricsContext';
import { formatDate } from '../../utils/formatUtils';

export const PAGE = 'page';
export const ROWS_PER_PAGE = 'rows';
export const FILTER = 'filter';
export const DATE = 'date';
export const PERIOD = 'period';
export const SORT_BY = 'sort';
export const ASCENDING = 'asc';
export const TIMELINE_VIEW = 'tl';
export const DIRECTORY_VIEW = 'dir';
export const SORT_INDEX = 'ind';

function ResourcesParamControls() {
  const { params } = useContext(MetricsContext);

  const [search, setSearchParams] = useSearchParams();

  const updateParams = useCallback(() => {
    search.set(PAGE, String(params.page));
    search.set(ROWS_PER_PAGE, String(params.rowsPerPage));
    search.set(FILTER, params.filter);
    search.set(DATE, formatDate(params.date));
    search.set(PERIOD, String(params.period));
    search.set(SORT_BY, String(params.sort));
    search.set(ASCENDING, String(params.ascending));
    search.set(TIMELINE_VIEW, String(params.timelineView));
    search.set(DIRECTORY_VIEW, String(params.directoryView));
    search.set(SORT_INDEX, String(params.sortIndex));
    setSearchParams(search);
  }, [search, setSearchParams, params]);

  useEffect(() => {
    updateParams();
  }, [updateParams]);

  return (<></>);
}

export default ResourcesParamControls;
