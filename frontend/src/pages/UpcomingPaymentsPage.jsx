import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import apiClient from '../api/client'
import { formatDate, getDaysLeft } from '../utils/date'

function UpcomingPaymentsPage() {
  const navigate = useNavigate()

  const [items, setItems] = useState([])
  const [days, setDays] = useState(30)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  async function fetchItems(targetDays = days) {
    setLoading(true)
    setError('')

    try {
      const response = await apiClient.get(`/figures/upcoming-payments?days=${targetDays}`)
      setItems(response.data.items || [])
      setDays(response.data.days || targetDays)
    } catch (err) {
      setError('無法取得即將補款資料')
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    const token = localStorage.getItem('token')

    if (!token) {
      navigate('/login')
      return
    }

    fetchItems(30)
  }, [])

  return (
    <div className="min-h-screen bg-slate-100">
      <header className="border-b border-slate-200 bg-white">
        <div className="mx-auto flex max-w-6xl items-center justify-between px-6 py-5">
          <div>
            <h1 className="text-2xl font-bold text-slate-900">收藏櫃</h1>
            <p className="text-sm text-slate-500">即將補款</p>
          </div>

          <div className="flex items-center gap-3">
            <button
              onClick={() => navigate('/dashboard')}
              className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
              總覽
            </button>

            <button
              onClick={() => navigate('/figures')}
              className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
              收藏列表
            </button>
          </div>
        </div>
      </header>

      <main className="mx-auto max-w-6xl px-6 py-8">
        <div className="mb-6 flex flex-col justify-between gap-4 md:flex-row md:items-end">
          <div>
            <h2 className="text-xl font-bold text-slate-900">即將補款</h2>
            <p className="mt-1 text-sm text-slate-500">
              查看未來 {days} 天內需要補款的公仔
            </p>
          </div>

          <div className="flex items-center gap-3">
            <select
              value={days}
              onChange={(event) => fetchItems(Number(event.target.value))}
              className="rounded-xl border border-slate-300 bg-white px-4 py-2 text-sm outline-none focus:border-slate-900"
            >
              <option value={7}>未來 7 天</option>
              <option value={30}>未來 30 天</option>
              <option value={60}>未來 60 天</option>
              <option value={365}>未來 365 天</option>
            </select>
          </div>
        </div>

        {error && (
          <div className="mb-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600">
            {error}
          </div>
        )}

        {loading ? (
          <div className="rounded-2xl bg-white p-10 text-center shadow-sm">
            <p className="text-slate-600">載入中...</p>
          </div>
        ) : items.length === 0 ? (
          <div className="rounded-2xl bg-white p-10 text-center shadow-sm">
            <p className="text-slate-600">目前沒有即將補款的公仔</p>
          </div>
        ) : (
          <div className="overflow-hidden rounded-2xl bg-white shadow-sm">
            <table className="w-full border-collapse text-left">
              <thead className="bg-slate-50 text-sm text-slate-500">
                <tr>
                  <th className="px-5 py-4 font-medium">名稱</th>
                  <th className="px-5 py-4 font-medium">角色</th>
                  <th className="px-5 py-4 font-medium">狀態</th>
                  <th className="px-5 py-4 font-medium">補款日</th>
                  <th className="px-5 py-4 font-medium">剩餘天數</th>
                  <th className="px-5 py-4 font-medium">尾款</th>
                  <th className="px-5 py-4 font-medium">平台</th>
                  <th className="px-5 py-4 font-medium">操作</th>
                </tr>
              </thead>

              <tbody className="divide-y divide-slate-100 text-sm">
                {items.map((figure) => (
                  <tr key={figure.id} className="hover:bg-slate-50">
                    <td className="px-5 py-4 font-medium text-slate-900">
                      {figure.name}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {figure.character_name || '-'}
                    </td>
                    <td className="px-5 py-4">
                      <span className="rounded-full bg-amber-100 px-3 py-1 text-xs text-amber-700">
                        {figure.status}
                      </span>
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {formatDate(figure.payment_due_date)}
                    </td>
                    <td className="px-5 py-4 font-medium text-slate-900">
                      {getDaysLeft(figure.payment_due_date)}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      NT$ {figure.balance}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {figure.shop_name || '-'}
                    </td>
                    <td className="px-5 py-4">
                      <button
                        onClick={() => navigate(`/figures/${figure.id}`)}
                        className="rounded-lg bg-slate-900 px-3 py-2 text-xs font-medium text-white hover:bg-slate-700"
                      >
                        查看
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </main>
    </div>
  )
}

export default UpcomingPaymentsPage