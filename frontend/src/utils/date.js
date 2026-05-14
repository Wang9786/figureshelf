export function formatDate(value) {
  if (!value) return '-'
  return value.slice(0, 10)
}

export function getDaysLeft(value) {
  if (!value) return '-'

  const targetDate = new Date(value.slice(0, 10))
  const today = new Date()

  targetDate.setHours(0, 0, 0, 0)
  today.setHours(0, 0, 0, 0)

  const diffMs = targetDate.getTime() - today.getTime()
  const diffDays = Math.ceil(diffMs / (1000 * 60 * 60 * 24))

  if (diffDays < 0) return `已過期 ${Math.abs(diffDays)} 天`
  if (diffDays === 0) return '今天'
  return `剩 ${diffDays} 天`
}