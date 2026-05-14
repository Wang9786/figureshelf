import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import apiClient from '../api/client'

function LoginPage() {
  const navigate = useNavigate()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(event) {
    event.preventDefault()
    setError('')
    setLoading(true)

    try {
      const response = await apiClient.post('/auth/login', {
        email,
        password,
      })

      localStorage.setItem('token', response.data.token)
      localStorage.setItem('user', JSON.stringify(response.data.user))

      navigate('/dashboard')
    } catch (err) {
      setError('登入失敗，請確認 Email 或密碼是否正確')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-100 px-4">
      <div className="w-full max-w-md rounded-2xl bg-white p-8 shadow-sm">
        <h1 className="text-3xl font-bold text-slate-900">收藏櫃</h1>
        <p className="mt-2 text-slate-500">登入你的公仔收藏管理系統</p>

        <form onSubmit={handleSubmit} className="mt-8 space-y-5">
          <div>
            <label className="block text-sm font-medium text-slate-700">
              Email
            </label>
            <input
              type="email"
              className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 outline-none focus:border-slate-900"
              value={email}
              onChange={(event) => setEmail(event.target.value)}
              placeholder="請輸入email"
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-700">
              密碼
            </label>
            <input
              type="password"
              className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 outline-none focus:border-slate-900"
              value={password}
              onChange={(event) => setPassword(event.target.value)}
              placeholder="請輸入密碼"
            />
          </div>

          {error && (
            <div className="rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600">
              {error}
            </div>
          )}

          <button
            type="submit"
            disabled={loading}
            className="w-full rounded-xl bg-slate-900 px-4 py-3 font-medium text-white hover:bg-slate-700 disabled:cursor-not-allowed disabled:bg-slate-400"
          >
            {loading ? '登入中...' : '登入'}
          </button>
        </form>

        <div className="mt-6 space-y-3 text-center text-sm text-slate-500">
          <p>
            還沒有帳號？
            <button
              onClick={() => navigate('/register')}
              className="ml-1 font-medium text-slate-900 hover:underline"
            >
              前往註冊
            </button>
          </p>
        </div>
      </div>
    </div>
  )
}

export default LoginPage