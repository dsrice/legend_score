import React, { useState } from 'react';
import {Dialog, DialogPanel, DialogTitle} from '@headlessui/react';
import { apiPost } from '../services/apiClient';
import { useTranslation } from 'react-i18next';

// Define the CreateUserRequest interface
interface CreateUserRequest {
  login_id: string;
  name: string;
  password: string;
}

// Define the CreateUserResponse interface
interface CreateUserResponse {
  result: boolean;
  message?: string;
}

interface CreateUserDialogProps {
  isOpen: boolean;
  onClose: () => void;
  onUserCreated: () => void;
}

const CreateUserDialog: React.FC<CreateUserDialogProps> = ({ isOpen, onClose, onUserCreated }) => {
  const { t } = useTranslation(); // Initialize translation function

  const [newUser, setNewUser] = useState<CreateUserRequest>({
    login_id: '',
    name: '',
    password: '',
  });
  const [createUserLoading, setCreateUserLoading] = useState<boolean>(false);
  const [createUserError, setCreateUserError] = useState<string | null>(null);

  // Reset form when dialog opens
  React.useEffect(() => {
    if (isOpen) {
      setNewUser({
        login_id: '',
        name: '',
        password: '',
      });
      setCreateUserError(null);
    }
  }, [isOpen]);

  // Handle input change for new user form
  const handleNewUserInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setNewUser(prev => ({
      ...prev,
      [name]: value
    }));
  };

  // Handle create user form submission
  const handleCreateUser = async (e: React.FormEvent) => {
    e.preventDefault();
    setCreateUserLoading(true);
    setCreateUserError(null);

    // Check if all required fields have values
    if (!newUser.login_id || !newUser.name || !newUser.password) {
      setCreateUserError(t('createUserDialog.error.allFieldsRequired') || 'All fields are required');
      setCreateUserLoading(false);
      return;
    }

    try {
      // Make the API request to create a user
      // Token is automatically added by the apiClient interceptor
      const response = await apiPost('/user', newUser);

      // Handle the response
      const data = response as CreateUserResponse;
      if (data.result) {
        // Close the dialog and notify parent component
        onClose();
        onUserCreated();
      } else {
        setCreateUserError(data.message || t('createUserDialog.error.createFailed'));
      }
    } catch (err: any) {
      console.error('Error creating user:', err);
      setCreateUserError(err.response?.data?.message || t('createUserDialog.error.generalError'));
    } finally {
      setCreateUserLoading(false);
    }
  };

  // Create a no-op function to prevent dialog from closing when clicking outside or pressing Escape
  const handleDialogClose = () => {
    // Do nothing, only allow closing via the Cancel button
  };

  return (
    <Dialog
      open={isOpen}
      onClose={handleDialogClose}
      className="fixed inset-0 z-10 overflow-y-auto"
    >
      <div className="flex items-center justify-center min-h-screen">
        <DialogPanel className="fixed inset-0 bg-black opacity-30" />

        <div className="relative bg-white rounded-lg max-w-md w-full mx-auto p-6 shadow-xl">
          <DialogTitle className="text-lg font-medium text-gray-900 mb-4">
            {t('createUserDialog.title')}
          </DialogTitle>

          {createUserError && (
            <div className="bg-red-50 border-l-4 border-red-400 p-4 mb-4" data-testid="error-message">
              <div className="flex">
                <div className="flex-shrink-0">
                  <svg className="h-5 w-5 text-red-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                  </svg>
                </div>
                <div className="ml-3">
                  <p className="text-sm text-red-700">{createUserError}</p>
                </div>
              </div>
            </div>
          )}

          <form onSubmit={handleCreateUser}>
            <div className="space-y-4">
              <div>
                <label htmlFor="login_id" className="block text-sm font-medium text-gray-700">
                  {t('createUserDialog.loginId')}
                </label>
                <input
                  type="text"
                  name="login_id"
                  id="login_id"
                  required
                  value={newUser.login_id}
                  onChange={handleNewUserInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
              <div>
                <label htmlFor="name" className="block text-sm font-medium text-gray-700">
                  {t('createUserDialog.name')}
                </label>
                <input
                  type="text"
                  name="name"
                  id="name"
                  required
                  value={newUser.name}
                  onChange={handleNewUserInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
              <div>
                <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                  {t('createUserDialog.password')}
                </label>
                <input
                  type="password"
                  name="password"
                  id="password"
                  required
                  value={newUser.password}
                  onChange={handleNewUserInputChange}
                  className="mt-1 block w-full border border-gray-300 rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
                />
              </div>
            </div>

            <div className="mt-6 flex justify-end space-x-3">
              <button
                type="button"
                onClick={onClose}
                className="py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                disabled={createUserLoading}
              >
                {t('createUserDialog.cancel')}
              </button>
              <button
                type="submit"
                className="py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
                disabled={createUserLoading}
              >
                {createUserLoading ? t('createUserDialog.creating') : t('createUserDialog.create')}
              </button>
            </div>
          </form>
        </div>
      </div>
    </Dialog>
  );
};

export default CreateUserDialog;