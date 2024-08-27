/* eslint-disable react/no-unescaped-entities */
import { Link, useRouteError } from 'react-router-dom'

function ErrorBoundary(error) {
  const routerError = useRouteError()

  if (routerError || error) {
    return (
      <div
        style={{
          textAlign: 'center',
          padding: '15px'
        }}>
        <h1>Oops! Something went wrong</h1>
        <div>
          The page you are looking for doesn't exit or you don't have permission to see <br />
          <Link to="/">Go back to Home page</Link>
        </div>
      </div>
    )
  } else {
    // bubble to the nearest parent route
    throw error
  }
}

export default ErrorBoundary
