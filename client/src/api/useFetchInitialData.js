import {useCallback, useEffect, useState} from 'react';

import {fetchImpl} from './baseAPI';

export const useFetchInitialData = (apiPath) => {
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [response, setResponse] = useState(null);
  const fetchData = async () => {
    setIsLoading(true);
    try {
      const body = await fetchImpl(apiPath);
      setResponse(body);
      setIsLoading(false);
    } catch (error) {
      setError(error);
    }
  };
  useEffect(
    () => {
      fetchData();
    },
    [],
  );
  return {
    error,
    isLoading,
    response,
  };
};
