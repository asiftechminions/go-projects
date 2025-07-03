Introduces a circuit breaker using the gobreaker library to handle
failures in an unreliable external service with configurable thresholds
and timeout. Implements retry logic with exponential backoff to
increase fault tolerance. Adds an HTTP endpoint to trigger the
service call and manage responses based on the circuit breaker state.
