import {useEffect, useState} from 'react';

import {API_URI} from './baseAPI';

export const useBaseFetchInitialData = (apiPath) => {
  const [response, setResponse] = useState(null);
  const [error, setError] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  useEffect(
    () => {
      const fetchData = async () => {
        setIsLoading(true);
        try {
          const httpResponse = await fetch(
            API_URI + apiPath,
            {
              credentials: 'same-origin',
              method: 'GET',
              mode: 'no-cors',
            },
          );
          const json = await httpResponse.json();
          setResponse(json);
          setIsLoading(false);
        } catch (error) {
          setError(error);
        }
      };
      fetchData();
    },
    [],
  );
  return {error, isLoading, response};
};
