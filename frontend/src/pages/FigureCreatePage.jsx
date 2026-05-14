import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import apiClient from '../api/client'

const initialForm = {
  name: '',
  character_name: '',
  series_name: '',
  manufacturer: '',
  figure_type: 'PVC',
  scale: '',
  status: 'wishlist',
  price: '',
  deposit: '',
  balance: '',
  preorder_start_date: '',
  preorder_deadline: '',
  release_date: '',
  payment_due_date: '',
  shop_name: '',
  note: '',
}

function FigureCreatePage() {
  const navigate = useNavigate()

  const [form, setForm] = useState(initialForm)
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  function handleChange(event) {
    const { name, value } = event.target

    setForm((prev) => ({
      ...prev,
      [name]: value,
    }))
  }

  function emptyToNull(value) {
    return value === '' ? null : value
  }

  async function handleSubmit(event) {
    event.preventDefault()
    setError('')
    setLoading(true)

    try {
      const payload = {
        name: form.name,
        character_name: emptyToNull(form.character_name),
        series_name: emptyToNull(form.series_name),
        manufacturer: emptyToNull(form.manufacturer),
        figure_type: emptyToNull(form.figure_type),
        scale: emptyToNull(form.scale),
        status: form.status,
        price: Number(form.price || 0),
        deposit: Number(form.deposit || 0),
        balance: Number(form.balance || 0),
        preorder_start_date: emptyToNull(form.preorder_start_date),
        preorder_deadline: emptyToNull(form.preorder_deadline),
        release_date: emptyToNull(form.release_date),
        payment_due_date: emptyToNull(form.payment_due_date),
        shop_name: emptyToNull(form.shop_name),
        note: emptyToNull(form.note),
      }

      await apiClient.post('/figures', payload)
      navigate('/figures')
    } catch (err) {
      setError('新增失敗，請確認必填欄位與日期格式')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-slate-100">
      <header className="border-b border-slate-200 bg-white">
        <div className="mx-auto flex max-w-6xl items-center justify-between px-6 py-5">
          <div>
            <h1 className="text-2xl font-bold text-slate-900">收藏櫃</h1>
            <p className="text-sm text-slate-500">新增公仔收藏</p>
          </div>

          <button
            onClick={() => navigate('/figures')}
            className="rounded-xl border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 hover:bg-slate-100"
          >
            返回列表
          </button>
        </div>
      </header>

      <main className="mx-auto max-w-4xl px-6 py-8">
        <div className="rounded-2xl bg-white p-8 shadow-sm">
          <h2 className="text-xl font-bold text-slate-900">新增收藏資料</h2>
          <p className="mt-1 text-sm text-slate-500">
            記錄公仔名稱、預購狀態、價格、補款日與發售日
          </p>

          {error && (
            <div className="mt-6 rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="mt-8 space-y-8">
            <section>
              <h3 className="mb-4 font-semibold text-slate-900">基本資料</h3>

              <div className="grid gap-5 md:grid-cols-2">
                <Input
                  label="公仔名稱 *"
                  name="name"
                  value={form.name}
                  onChange={handleChange}
                  required
                />

                <Input
                  label="角色名稱"
                  name="character_name"
                  value={form.character_name}
                  onChange={handleChange}
                />

                <Input
                  label="作品名稱"
                  name="series_name"
                  value={form.series_name}
                  onChange={handleChange}
                />

                <Input
                  label="廠商 / 工作室"
                  name="manufacturer"
                  value={form.manufacturer}
                  onChange={handleChange}
                />

                <Input
                  label="類型"
                  name="figure_type"
                  value={form.figure_type}
                  onChange={handleChange}
                />

                <Input
                  label="比例"
                  name="scale"
                  value={form.scale}
                  onChange={handleChange}
                  placeholder="例如：1/7"
                />
              </div>
            </section>

            <section>
              <h3 className="mb-4 font-semibold text-slate-900">狀態與金額</h3>

              <div className="grid gap-5 md:grid-cols-2">
                <div>
                  <label className="block text-sm font-medium text-slate-700">
                    狀態
                  </label>
                  <select
                    name="status"
                    value={form.status}
                    onChange={handleChange}
                    className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 outline-none focus:border-slate-900"
                  >
                    <option value="wishlist">想買</option>
                    <option value="preordered">已預購</option>
                    <option value="deposit_paid">已付訂金</option>
                    <option value="balance_due">待補款</option>
                    <option value="paid">已付清</option>
                    <option value="shipped">已出貨</option>
                    <option value="arrived">已到貨</option>
                    <option value="cancelled">已取消</option>
                    <option value="sold">已售出</option>
                  </select>
                </div>

                <Input
                  label="購入價格"
                  name="price"
                  type="number"
                  value={form.price}
                  onChange={handleChange}
                />

                <Input
                  label="訂金"
                  name="deposit"
                  type="number"
                  value={form.deposit}
                  onChange={handleChange}
                />

                <Input
                  label="尾款"
                  name="balance"
                  type="number"
                  value={form.balance}
                  onChange={handleChange}
                />
              </div>
            </section>

            <section>
              <h3 className="mb-4 font-semibold text-slate-900">日期資訊</h3>

              <div className="grid gap-5 md:grid-cols-2">
                <Input
                  label="開訂日"
                  name="preorder_start_date"
                  type="date"
                  value={form.preorder_start_date}
                  onChange={handleChange}
                />

                <Input
                  label="截單日"
                  name="preorder_deadline"
                  type="date"
                  value={form.preorder_deadline}
                  onChange={handleChange}
                />

                <Input
                  label="預計發售日"
                  name="release_date"
                  type="date"
                  value={form.release_date}
                  onChange={handleChange}
                />

                <Input
                  label="補款期限"
                  name="payment_due_date"
                  type="date"
                  value={form.payment_due_date}
                  onChange={handleChange}
                />
              </div>
            </section>

            <section>
              <h3 className="mb-4 font-semibold text-slate-900">其他資訊</h3>

              <div className="grid gap-5">
                <Input
                  label="購買平台"
                  name="shop_name"
                  value={form.shop_name}
                  onChange={handleChange}
                  placeholder="例如：AmiAmi、露天、蝦皮、工作室"
                />

                <div>
                  <label className="block text-sm font-medium text-slate-700">
                    備註
                  </label>
                  <textarea
                    name="note"
                    value={form.note}
                    onChange={handleChange}
                    rows="4"
                    className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 outline-none focus:border-slate-900"
                  />
                </div>
              </div>
            </section>

            <div className="flex justify-end gap-3 border-t border-slate-100 pt-6">
              <button
                type="button"
                onClick={() => navigate('/figures')}
                className="rounded-xl border border-slate-300 px-5 py-3 text-sm font-medium text-slate-700 hover:bg-slate-100"
              >
                取消
              </button>

              <button
                type="submit"
                disabled={loading}
                className="rounded-xl bg-slate-900 px-5 py-3 text-sm font-medium text-white hover:bg-slate-700 disabled:cursor-not-allowed disabled:bg-slate-400"
              >
                {loading ? '儲存中...' : '儲存'}
              </button>
            </div>
          </form>
        </div>
      </main>
    </div>
  )
}

function Input({
  label,
  name,
  value,
  onChange,
  type = 'text',
  required = false,
  placeholder = '',
}) {
  return (
    <div>
      <label className="block text-sm font-medium text-slate-700">
        {label}
      </label>
      <input
        type={type}
        name={name}
        value={value}
        onChange={onChange}
        required={required}
        placeholder={placeholder}
        className="mt-2 w-full rounded-xl border border-slate-300 px-4 py-3 outline-none focus:border-slate-900"
      />
    </div>
  )
}

export default FigureCreatePage