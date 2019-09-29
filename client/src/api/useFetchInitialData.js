// @flow strict

import {useCallback, useEffect, useState} from 'react';

import {fetchImpl} from './baseAPI';

type InitialData<T> = {
  data: ?T,
  error: ?Error,
  isLoading: boolean,
};

export function useFetchInitialData<T>(
  apiPath: string
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
    },
    [],
  );
  useEffect(
    (): void => {
      fetchData();
    },
    [],
  );
  return {
    data,
    error,
    isLoading,
  };
};
