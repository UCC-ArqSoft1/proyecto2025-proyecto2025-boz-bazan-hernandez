// App.js
import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import Home from './pages/Home/Home';
import Activities from './pages/Activities/Activities';
import ActivityDetailsPage from './pages/Activities/ActivityDetailsPage';
import MyActivitiesPage from './pages/Activities/MyActivitiesPage';
import LoginPage from './pages/Auth/LoginPage';
import RegisterPage from './pages/Auth/RegisterPage';
import AdminActivitiesPage from './pages/Admin/AdminActivitiesPage';
import CreateActivityPage from './pages/Admin/CreateActivityPage';
import EditActivityPage from './pages/Admin/EditActivityPage';
import NotFoundPage from './pages/NotFoundPage';
import ProtectedRoute from './components/Auth/ProtectedRoute';
import AdminRoute from './components/Auth/AdminRoute';
import Layout from './components/Layout/Layout';

function App() {
  return (
    <Router>
      <AuthProvider>
        <Layout>
          <Switch>
            <Route exact path="/" component={Home} />
            <Route exact path="/activities" component={Activities} />
            <Route exact path="/activities/:id" component={ActivityDetailsPage} />
            <ProtectedRoute exact path="/my-activities" component={MyActivitiesPage} />
            <Route exact path="/login" component={LoginPage} />
            <Route exact path="/register" component={RegisterPage} />
            <AdminRoute exact path="/admin/activities" component={AdminActivitiesPage} />
            <AdminRoute exact path="/admin/activities/create" component={CreateActivityPage} />
            <AdminRoute exact path="/admin/activities/edit/:id" component={EditActivityPage} />
            <Route component={NotFoundPage} />
          </Switch>
        </Layout>
      </AuthProvider>
    </Router>
  );
}

export default App;