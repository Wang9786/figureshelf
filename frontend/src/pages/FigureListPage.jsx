import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import apiClient from '../api/client'

function FigureListPage() {
  const navigate = useNavigate()

  const [figures, setFigures] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  async function fetchFigures() {
    try {
      const response = await apiClient.get('/figures')
      setFigures(response.data.items || [])
    } catch (err) {
      setError('無法取得收藏列表，請重新登入')
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

    fetchFigures()
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
            <p className="text-sm text-slate-500">收藏列表</p>
          </div>

          <div className="flex items-center gap-3">
            <button
              onClick={() => navigate('/dashboard')}
              className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
              總覽
            </button>

            <button
              onClick={() => navigate('/figures/new')}
              className="rounded-xl bg-slate-900 px-4 py-2 text-sm font-medium text-white hover:bg-slate-700"
            >
              新增公仔
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
        <div className="mb-6">
          <h2 className="text-xl font-bold text-slate-900">我的收藏</h2>
          <p className="mt-1 text-sm text-slate-500">
            查看所有已建立的公仔收藏與預購資料
          </p>
        </div>

        {error && (
          <div className="mb-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600">
            {error}
          </div>
        )}

        {figures.length === 0 ? (
          <div className="rounded-2xl bg-white p-10 text-center shadow-sm">
            <p className="text-slate-600">目前還沒有收藏資料</p>
            <button
              onClick={() => navigate('/figures/new')}
              className="mt-5 rounded-xl bg-slate-900 px-5 py-3 text-sm font-medium text-white hover:bg-slate-700"
            >
              新增第一筆公仔
            </button>
          </div>
        ) : (
          <div className="overflow-hidden rounded-2xl bg-white shadow-sm">
            <table className="w-full border-collapse text-left">
              <thead className="bg-slate-50 text-sm text-slate-500">
                <tr>
                    <th className="px-5 py-4 font-medium">名稱</th>
                    <th className="px-5 py-4 font-medium">角色</th>
                    <th className="px-5 py-4 font-medium">作品</th>
                    <th className="px-5 py-4 font-medium">狀態</th>
                    <th className="px-5 py-4 font-medium">價格</th>
                    <th className="px-5 py-4 font-medium">補款日</th>
                    <th className="px-5 py-4 font-medium">發售日</th>
                    <th className="px-5 py-4 font-medium">操作</th>
                </tr>
              </thead>

              <tbody className="divide-y divide-slate-100 text-sm">
                {figures.map((figure) => (
                  <tr key={figure.id} className="hover:bg-slate-50">
                    <td className="px-5 py-4 font-medium text-slate-900">
                      {figure.name}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {figure.character_name || '-'}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {figure.series_name || '-'}
                    </td>
                    <td className="px-5 py-4">
                      <span className="rounded-full bg-slate-100 px-3 py-1 text-xs text-slate-700">
                        {figure.status}
                      </span>
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      NT$ {figure.price}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {formatDate(figure.payment_due_date)}
                    </td>
                    <td className="px-5 py-4 text-slate-600">
                      {formatDate(figure.release_date)}
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

function formatDate(value) {
  if (!value) return '-'
  return value.slice(0, 10)
}

export default FigureListPage