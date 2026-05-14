import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import apiClient from '../api/client'

function RegisterPage() {
  const navigate = useNavigate()

  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [confirmPassword, setConfirmPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  async function handleSubmit(event) {
    event.preventDefault()
    setError('')

    if (password.length < 6) {
      setError('密碼至少需要 6 個字元')
      return
    }

    if (password !== confirmPassword) {
      setError('兩次輸入的密碼不一致')
      return
    }

    setLoading(true)

    try {
      await apiClient.post('/auth/register', {
        email,
        password,
      })

      navigate('/login')
    } catch (err) {
      if (err.response?.status === 409) {
        setError('這個 Email 已經註冊過了')
      } else {
        setError('註冊失敗，請稍後再試')
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="flex min-h-screen items-center justify-center bg-slate-100 px-4">
      <div className="w-full max-w-md rounded-2xl bg-white p-8 shadow-sm">
        <h1 className="text-3xl font-bold text-slate-900">收藏櫃</h1>
        <p className="mt-2 text-slate-500">建立你的公仔收藏管理帳號</p>

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
              placeholder="you@example.com"
              required
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
              placeholder="至少 6 個字元"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-700">
              確認密碼
            </label>
            <input
              type="password"
              className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 outline-none focus:border-slate-900"
              value={confirmPassword}
              onChange={(event) => setConfirmPassword(event.target.value)}
              placeholder="再次輸入密碼"
              required
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
            {loading ? '註冊中...' : '註冊'}
          </button>
        </form>

        <div className="mt-6 text-center text-sm text-slate-500">
          已經有帳號？
          <button
            onClick={() => navigate('/login')}
            className="ml-1 font-medium text-slate-900 hover:underline"
          >
            前往登入
          </button>
        </div>
      </div>
    </div>
  )
}

export default RegisterPage