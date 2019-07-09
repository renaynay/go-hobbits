package encoding_test

//func TestUnmarshal_Successful(t *testing.T) {
//	var test = []struct {
//		message string
//		output  encoding.Message
//	}{
//		{
//			message: "EWP 13.05 RPC 16 14\nthis is a headerthis is a body",
//			output: encoding.Message{
//				Version:  "13.05",
//				Protocol: encoding.RPC,
//				Header:   []byte("this is a header"),
//				Body:     []byte("this is a body"),
//			},
//		},
//		{
//			message: "EWP 13.05 GOSSIP 7 12\ntestingtesting body",
//			output: encoding.Message{
//				Version:  "13.05",
//				Protocol: encoding.GOSSIP,
//				Header:   []byte("testing"),
//				Body:     []byte("testing body"),
//			},
//		},
//		{
//			message: "EWP 1230329483.05392489 RPC 4 4\ntesttest",
//			output: encoding.Message{
//				Version:  "1230329483.05392489",
//				Protocol: encoding.RPC,
//				Header:   []byte("test"),
//				Body:     []byte("test"),
//			},
//		},
//		{
//			message: "EWP 1230329483.05392489 PING 4 4\ntesttest",
//			output: encoding.Message{
//				Version:  "1230329483.05392489",
//				Protocol: "PING",
//				Header:   []byte("test"),
//				Body:     []byte("test"),
//			},
//		},
//	}
//
//	for i, tt := range test {
//		t.Run(strconv.Itoa(i), func(t *testing.T) {
//			output, _ := encoding.Unmarshal(tt.message)
//			if !reflect.DeepEqual(*output, tt.output) {
//				t.Errorf("return value of Unmarshal does not match expected value")
//			}
//		})
//	}
//}
//
//func TestUnmarshal_Unsuccessful(t *testing.T) {
//	var test = []struct {
//		message string
//		err     error
//	}{
//		{
//			message: "EWP 13.05 RPC blahblahblah json 16 14this is a headerthis is a body",
//			err:     errors.New("message request must contain 2 lines"),
//		},
//		{
//			message: "EWP 13.05 7 12\ntestingtesting body",
//			err:     errors.New("not all metadata provided"),
//		},
//		{
//			message: "EWP 123032948392489 RPC 4 4\ntesttest",
//			err:     errors.New("EWP version cannot be parsed"),
//		},
//		{
//			message: "EWP 123032948.392489 notrpc 4 4\ntesttest",
//			err:     errors.New("communication protocol unsupported"),
//		},
//		{
//			message: "EWP 123032948.392489 GOSSIP f 4\ntesttest",
//			err:     errors.New("incorrect metadata format, cannot parse header-length"),
//		},
//		{
//			message: "EWP 123032948.392489 GOSSIP 4 f\ntesttest",
//			err:     errors.New("incorrect metadata format, cannot parse body-length"),
//		},
//	}
//
//	for i, tt := range test {
//		t.Run(strconv.Itoa(i), func(t *testing.T) {
//			_, err := encoding.Unmarshal(tt.message)
//			if !reflect.DeepEqual(err, tt.err) {
//				t.Errorf("return value of Unmarshal did not match expected value")
//			}
//		})
//	}
//}
//
//func BenchmarkUnmarshal(b *testing.B) {
//	for i := 0; i <= b.N; i++ {
//		encoding.Unmarshal("EWP 13.05 RPC 16 14\nthis is a headerthis is a body")
//	}
//}
