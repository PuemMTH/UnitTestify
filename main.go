package main

import (
	"fmt"
	"net/http"
	"time"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*") // หรือระบุโดเมนเฉพาะ
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func main() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// ตั้งค่า CORS
		enableCors(&w)

		// ตั้งค่าสำหรับ SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// it protocol for SSE
		count := 0
		for {
			count++
			fmt.Fprint(w, "data: "+time.Now().Format(time.RFC1123)+"\t"+fmt.Sprint(count)+"\n\n")
			w.(http.Flusher).Flush()
			time.Sleep(500 * time.Millisecond)
		}
	})

	// เริ่มเซิร์ฟเวอร์
	http.ListenAndServe(":8080", nil)
}

// const [data, setData] = useState("");
// useEffect(() => {
// 	// สร้างเชื่อมต่อ SSE ไปยังเซิร์ฟเวอร์ Go
// 	const eventSource = new EventSource('http://localhost:8080/events');

// 	// รับข้อมูลจากเซิร์ฟเวอร์
// 	eventSource.onmessage = (e) => {
// 		setData(e.data);
// 	};

// 	// ปิดเชื่อมต่อเมื่อออกจากหน้าเพจ
// 	return () => {
// 		eventSource.close();
// 	};
// }, []);
