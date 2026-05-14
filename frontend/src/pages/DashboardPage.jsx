import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import apiClient from '../api/client'

function StatCard({ title, value }) {
  return (
    <div className="rounded-2xl border border-slate-200 bg-white p-6 shadow-sm">
      <p className="text-sm text-slate-500">{title}</p>
      <p className="mt-3 text-3xl font-bold text-slate-900">{value}</p>
    </div>
  )
}

function DashboardPage() {
  const navigate = useNavigate()

  const [summary, setSummary] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  async function fetchSummary() {
    try {
      const response = await apiClient.get('/dashboard/summary')
      setSummary(response.data)
    } catch (err) {
      setError('無法取得 Dashboard 資料，請重新登入')
    } finally {
      setLoading(false)
    }
  }

  function handleLogout() {
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    navigate('/login')
  }

  useEffect(() => {
    const token = localStorage.getItem('token')

    if (!token) {
      navigate('/login')
      return
    }

    fetchSummary()
  }, [])

  if (loading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-slate-100">
        <p className="text-slate-600">載入中...</p>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-slate-100">
      <header className="border-b border-slate-200 bg-white">
        <div className="mx-auto flex max-w-6xl items-center justify-between px-6 py-5">
          <div>
            <h1 className="text-2xl font-bold text-slate-900">收藏櫃</h1>
            <p className="text-sm text-slate-500">公仔收藏與預購管理系統</p>
          </div>

          <div className="flex items-center gap-3">
            <button
                onClick={() => navigate('/figures')}
                className="rounded-xl bg-slate-900 px-4 py-2 text-sm font-medium text-white hover:bg-slate-700"
            >
                收藏列表
            </button>

            <button
                onClick={handleLogout}
                className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
                登出
            </button>
            </div>
        </div>
      </header>

      <main className="mx-auto max-w-6xl px-6 py-8">
        {error && (
          <div className="mb-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600">
            {error}
          </div>
        )}

        <div className="mb-6 flex items-end justify-between">
          <div>
            <h2 className="text-xl font-bold text-slate-900">總覽</h2>
            <p className="mt-1 text-sm text-slate-500">
              查看你的收藏數量、預購狀態與補款資訊
            </p>
          </div>

          <div className="flex items-center gap-3">
            <button
                onClick={() => navigate('/figures/upcoming-payments')}
                className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
                即將補款
            </button>

            <button
                onClick={() => navigate('/figures/upcoming-releases')}
                className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
                即將發售
            </button>
            </div>
        </div>

        <div className="grid gap-5 md:grid-cols-3">
          <StatCard title="收藏總數" value={summary?.total_figures ?? 0} />
          <StatCard title="願望清單" value={summary?.wishlist_count ?? 0} />
          <StatCard title="預購中" value={summary?.preordered_count ?? 0} />
          <StatCard title="已到貨" value={summary?.arrived_count ?? 0} />
          <StatCard title="即將補款" value={summary?.upcoming_payments_count ?? 0} />
          <StatCard title="即將發售" value={summary?.upcoming_releases_count ?? 0} />
        </div>

        <div className="mt-8 grid gap-5 md:grid-cols-3">
          <StatCard title="收藏總金額" value={`NT$ ${summary?.total_price ?? 0}`} />
          <StatCard title="已付訂金" value={`NT$ ${summary?.total_deposit ?? 0}`} />
          <StatCard title="待付尾款" value={`NT$ ${summary?.total_balance ?? 0}`} />
        </div>
      </main>
    </div>
  )
}

export default DashboardPage