import { useState } from "react"

type OriginalUrlProps = { url?: string }

const OriginalUrl = ({ url }: OriginalUrlProps) => {
  const [truncate, setTruncate] = useState(true)

  const toggleTruncate = () => {
    setTruncate(!truncate)
  }
  return (
    <div className="mt-3">
      <div className="text-gray-600">Original URL</div>
      <div
        className={`cursor-pointer break-all ${truncate ? "truncate" : ""}`}
        onClick={toggleTruncate}
      >
        {url}
      </div>
    </div>
  )
}

export default OriginalUrl
