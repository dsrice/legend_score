import React, { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { apiGet } from '../services/apiClient';
import { getToken } from '../services/auth';

// Define the User interface based on the backend response
interface User {
  id: number;
  login_id: string;
  name: string;
}

// Define the API response interface
interface GetUsersResponse {
  result: boolean;
  users: User[];
}

const UserList: React.FC = () => {
  const navigate = useNavigate();
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // Filter state
  const [userIdFilter, setUserIdFilter] = useState<string>('');
  const [loginIdFilter, setLoginIdFilter] = useState<string>('');
  const [nameFilter, setNameFilter] = useState<string>('');

  // Fetch users on component mount
  useEffect(() => {
    fetchUsers();
  }, []);

  // Function to fetch users with optional filters
  const fetchUsers = async (filters: { user_id?: string; login_id?: string; name?: string } = {}) => {
    setLoading(true);
    setError(null);

    try {
      // Prepare query parameters
      const params = new URLSearchParams();
      if (filters.user_id) params.append('user_id', filters.user_id);
      if (filters.login_id) params.append('login_id', filters.login_id);
      if (filters.name) params.append('name', filters.name);

      // Get the token for authentication
      const token = getToken();

      // Make the API request
      const response = await apiGet(`/users?${params.toString()}`, {
        headers: {
          Authorization: `Bearer ${token}`
        }
      });

      // Update state with the response data
      const data = response as GetUsersResponse;
      if (data.result) {
        setUsers(data.users);
      } else {
        setError('Failed to fetch users');
      }
    } catch (err) {
      console.error('Error fetching users:', err);
      setError('An error occurred while fetching users');
    } finally {
      setLoading(false);
    }
  };

  // Handle filter submission
  const handleFilterSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    fetchUsers({
      user_id: userIdFilter,
      login_id: loginIdFilter,
      name: nameFilter
    });
  };

  // Handle filter reset
  const handleFilterReset = () => {
    setUserIdFilter('');
    setLoginIdFilter('');
    setNameFilter('');
    fetchUsers();
  };

  // Navigate back to home
  const handleBackToHome = () => {
    navigate('/home');
  };

  return (
    <div className="min-h-screen bg-gray-50 py-8 px-4 sm:px-6 lg:px-8">
      <div className="max-w-7xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold text-gray-900">User List</h1>
          <button
            onClick={handleBackToHome}
            className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
          >
            Back to Home
          </button>
        </div>

        {/* Filter Form */}
        <div className="bg-white shadow rounded-lg mb-6 p-4">
          <h2 className="text-lg font-medium mb-4">Filter Users</h2>
          <form onSubmit={handleFilterSubmit} className="grid grid-cols-1 gap-4 sm:grid-cols-3">
            <div>
              <label htmlFor="user-id" className="block text-sm font-medium text-gray-700">User ID</label>
              <input
                type="text"
                id="user-id"
                value={userIdFilter}
                onChange={(e) => setUserIdFilter(e.target.value)}
                className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              />
            </div>
            <div>
              <label htmlFor="login-id" className="block text-sm font-medium text-gray-700">Login ID</label>
              <input
                type="text"
                id="login-id"
                value={loginIdFilter}
                onChange={(e) => setLoginIdFilter(e.target.value)}
                className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              />
            </div>
            <div>
              <label htmlFor="name" className="block text-sm font-medium text-gray-700">Name</label>
              <input
                type="text"
                id="name"
                value={nameFilter}
                onChange={(e) => setNameFilter(e.target.value)}
                className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              />
            </div>
            <div className="sm:col-span-3 flex justify-end space-x-3">
              <button
                type="button"
                onClick={handleFilterReset}
                className="py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Reset
              </button>
              <button
                type="submit"
                className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
              >
                Apply Filters
              </button>
            </div>
          </form>
        </div>

        {/* Error Message */}
        {error && (
          <div className="bg-red-50 border-l-4 border-red-400 p-4 mb-6">
            <div className="flex">
              <div className="flex-shrink-0">
                <svg className="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                </svg>
              </div>
              <div className="ml-3">
                <p className="text-sm text-red-700">{error}</p>
              </div>
            </div>
          </div>
        )}

        {/* User Table */}
        <div className="bg-white shadow overflow-hidden rounded-lg">
          {loading ? (
            <div className="p-6 text-center">
              <p className="text-gray-500">Loading users...</p>
            </div>
          ) : users.length === 0 ? (
            <div className="p-6 text-center">
              <p className="text-gray-500">No users found</p>
            </div>
          ) : (
            <table className="min-w-full divide-y divide-gray-200">
              <thead className="bg-gray-50">
                <tr>
                  <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    ID
                  </th>
                  <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Login ID
                  </th>
                  <th scope="col" className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                    Name
                  </th>
                </tr>
              </thead>
              <tbody className="bg-white divide-y divide-gray-200">
                {users.map((user) => (
                  <tr key={user.id}>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      {user.id}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      {user.login_id}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                      {user.name}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>
      </div>
    </div>
  );
};

export default UserList;