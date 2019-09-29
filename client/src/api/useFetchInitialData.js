import {useCallback, useEffect, useState} from 'react';

import {fetchImpl} from './baseAPI';

export const useFetchInitialData = (apiPath) => {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const fetchData = useCallback(
    async () => {
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
    () => {
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
