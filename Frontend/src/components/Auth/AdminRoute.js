import React, { useContext } from 'react';
import { Route, Redirect } from 'react-router-dom';
import { AuthContext } from '../../context/AuthContext';

const AdminRoute = ({ component: Component, ...rest }) => {
  const { user, loading, isAdmin } = useContext(AuthContext);

  if (loading) {
    return null; // Or a loading spinner
  }

  return (
    <Route
      {...rest}
      render={(props) =>
        user && isAdmin() ? (
          <Component {...props} />
        ) : (
          <Redirect to={{ pathname: '/', state: { from: props.location } }} />
        )
      }
    />
  );
};

export default AdminRoute;