import { useEffect, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import apiClient from '../api/client'

function FigureDetailPage() {
  const navigate = useNavigate()
  const { id } = useParams()

  const [figure, setFigure] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [deleting, setDeleting] = useState(false)

  async function fetchFigure() {
    try {
      const response = await apiClient.get(`/figures/${id}`)
      setFigure(response.data)
    } catch (err) {
      setError('無法取得公仔資料')
    } finally {
      setLoading(false)
    }
  }

  async function handleDelete() {
    const confirmed = window.confirm('確定要刪除這筆公仔資料嗎？刪除後無法復原。')

    if (!confirmed) {
      return
    }

    setDeleting(true)

    try {
      await apiClient.delete(`/figures/${id}`)
      navigate('/figures')
    } catch (err) {
      setError('刪除失敗，請稍後再試')
      setDeleting(false)
    }
  }

  useEffect(() => {
    const token = localStorage.getItem('token')

    if (!token) {
      navigate('/login')
      return
    }

    fetchFigure()
  }, [id])

  if (loading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-slate-100">
        <p className="text-slate-600">載入中...</p>
      </div>
    )
  }

  if (error && !figure) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-slate-100 px-4">
        <div className="rounded-2xl bg-white p-8 text-center shadow-sm">
          <p className="text-red-600">{error}</p>
          <button
            onClick={() => navigate('/figures')}
            className="mt-5 rounded-xl bg-slate-900 px-5 py-3 text-sm font-medium text-white hover:bg-slate-700"
          >
            返回列表
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-slate-100">
      <header className="border-b border-slate-200 bg-white">
        <div className="mx-auto flex max-w-6xl items-center justify-between px-6 py-5">
          <div>
            <h1 className="text-2xl font-bold text-slate-900">收藏櫃</h1>
            <p className="text-sm text-slate-500">公仔詳細資料</p>
          </div>

          <div className="flex items-center gap-3">
            <button
              onClick={() => navigate('/figures')}
              className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
              返回列表
            </button>

            <button
              onClick={() => navigate('/dashboard')}
              className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
            >
              總覽
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

        <div className="rounded-2xl bg-white p-8 shadow-sm">
          <div className="flex flex-col justify-between gap-5 border-b border-slate-100 pb-6 md:flex-row md:items-start">
            <div>
              <div className="flex items-center gap-3">
                <h2 className="text-2xl font-bold text-slate-900">
                  {figure.name}
                </h2>

                <span className="rounded-full bg-slate-100 px-3 py-1 text-xs text-slate-700">
                  {figure.status}
                </span>
              </div>

              <p className="mt-2 text-sm text-slate-500">
                {figure.series_name || '未填寫作品'} / {figure.character_name || '未填寫角色'}
              </p>
            </div>

            <div className="flex gap-3">
              <button
                onClick={() => navigate(`/figures/${id}/edit`)}
                className="rounded-xl bg-slate-900 px-5 py-3 text-sm font-medium text-white hover:bg-slate-700"
              >
                編輯
              </button>

              <button
                onClick={handleDelete}
                disabled={deleting}
                className="rounded-xl bg-red-600 px-5 py-3 text-sm font-medium text-white hover:bg-red-500 disabled:cursor-not-allowed disabled:bg-red-300"
              >
                {deleting ? '刪除中...' : '刪除'}
              </button>
            </div>
          </div>

          <section className="mt-8">
            <h3 className="text-lg font-semibold text-slate-900">基本資料</h3>

            <div className="mt-5 grid gap-4 md:grid-cols-3">
              <InfoItem label="公仔名稱" value={figure.name} />
              <InfoItem label="角色名稱" value={figure.character_name} />
              <InfoItem label="作品名稱" value={figure.series_name} />
              <InfoItem label="廠商 / 工作室" value={figure.manufacturer} />
              <InfoItem label="類型" value={figure.figure_type} />
              <InfoItem label="比例" value={figure.scale} />
              <InfoItem label="購買平台" value={figure.shop_name} />
            </div>
          </section>

          <section className="mt-8">
            <h3 className="text-lg font-semibold text-slate-900">金額資訊</h3>

            <div className="mt-5 grid gap-4 md:grid-cols-3">
              <InfoItem label="購入價格" value={`NT$ ${figure.price}`} />
              <InfoItem label="訂金" value={`NT$ ${figure.deposit}`} />
              <InfoItem label="尾款" value={`NT$ ${figure.balance}`} />
            </div>
          </section>

          <section className="mt-8">
            <h3 className="text-lg font-semibold text-slate-900">日期資訊</h3>

            <div className="mt-5 grid gap-4 md:grid-cols-3">
              <InfoItem label="開訂日" value={formatDate(figure.preorder_start_date)} />
              <InfoItem label="截單日" value={formatDate(figure.preorder_deadline)} />
              <InfoItem label="預計發售日" value={formatDate(figure.release_date)} />
              <InfoItem label="補款期限" value={formatDate(figure.payment_due_date)} />
              <InfoItem label="到貨日" value={formatDate(figure.arrival_date)} />
            </div>
          </section>

          <section className="mt-8">
            <h3 className="text-lg font-semibold text-slate-900">備註</h3>

            <div className="mt-5 rounded-xl bg-slate-50 p-5 text-sm text-slate-700">
              {figure.note || '尚無備註'}
            </div>
          </section>
        </div>
      </main>
    </div>
  )
}

function InfoItem({ label, value }) {
  return (
    <div className="rounded-xl border border-slate-200 p-5">
      <p className="text-sm text-slate-500">{label}</p>
      <p className="mt-2 font-medium text-slate-900">{value || '-'}</p>
    </div>
  )
}

function formatDate(value) {
  if (!value) return '-'
  return value.slice(0, 10)
}

export default FigureDetailPage