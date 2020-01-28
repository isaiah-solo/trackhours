// @flow strict

import {useCallback, useEffect, useState} from 'react';

import {fetchImpl} from './baseAPI';

type InitialData<T> = {
  data: ?T,
  error: ?Error,
  isLoading: boolean,
};

export function useFetchInitialData<T>(
  apiPath: string,
  onFetchComplete?: () => void,
  inputs?: Array<mixed>
): InitialData<T> {
  const [data, setData] = useState<?T>(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const fetchData = useCallback(
    async (): Promise<void> => {
      setIsLoading(true);
      try {
        const data = await fetchImpl(apiPath);
        setData(data);
        setIsLoading(false);
      } catch (error) {
        setError(error);
      }
      if (onFetchComplete !== undefined) {
        onFetchComplete();
      }
    },
    [],
  );
  useEffect(
    (): void => {
      fetchData();
    },
    inputs !== undefined ? inputs : [],
  );
  return {
    data,
    error,
    isLoading,
  };
};
