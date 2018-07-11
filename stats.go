package main

import (
	"strings"
	"github.com/iancoleman/strcase"
	"strconv"
	"reflect"
	"bytes"
	"fmt"
)

type Stats struct {
	Pid         int
	StartTime   int
	CurrentTime int
	Uptime      int
	//totIdleTime         int
	//averageIdlePercent  float64
	//recentIdlePercent   float64
	//activeNetworkEvents int
	//timeAfterEpoll      float64
	//epollCalls          float64
	//epollIntr           int
	//PID                 string

	//>>>>>>connections>>>>>>	start

	//ActiveConnections int
	//activeDhConnections          int
	//OutboundConnections int
	//readyOutboundConnections     int
	//activeOutboundConnections    int
	//outboundConnectionsCreated   int
	//totalConnectFailures         int
	//inboundConnections           int
	//activeInboundConnections     int
	//inboundConnectionsAccepted   int
	//listeningConnections         int
	//unusedConnectionsClosed      int
	ReadyTargets int
	//allocatedTargets             int
	ActiveTargets int
	//inactiveTargets              int
	//freeTargets                  int
	//maxConnections               int
	//activeSpecialConnections     int
	//maxSpecialConnections        int
	//maxAcceptRate                int
	//curAcceptRateRemaining       float64
	//maxConnection                int
	//connGeneration               int
	//allocatedConnections         int
	//allocatedOutboundConnections int
	//allocatedInboundConnections  int
	//allocatedSocketConnections   int
	//tcpReadvCalls                int
	//tcpReadvIntr                 int
	//tcpReadvBytes                int
	//tcpWritevCalls               int
	//tcpWritevIntr                int
	//tcpWritevBytes               int
	//freeLaterSize                int
	//freeLaterTotal               int
	//acceptCallsFailed            int
	//acceptNonblockSetFailed      int
	//acceptConnectionLimitFailed  int
	//acceptRateLimitFailed        int
	//acceptInitAcceptedFailed     int

	// <<<<<<connections<<<<<<	end

	//>>>>>>raw_msg>>>>>>	start

	//rwmTotalMsgs     int
	//rwmTotalMsgParts int

	//<<<<<<raw_msg<<<<<<	end

	//>>>>>>raw_msg_buffer>>>>>>	start

	//totalUsedBuffersSize     int
	//totalUsedBuffers         int
	//allocatedBufferBytes     int
	//bufferChunkAllocOps      int
	//allocatedBufferChunks    int
	//maxAllocatedBufferChunks int
	//maxBufferChunks          int
	//maxAllocatedBufferBytes  int

	//<<<<<<raw_msg_buffer<<<<<<	end

	//>>>>>>tl_parse>>>>>>	start

	//rpcQueriesReceived int
	//rpcAnswersError    int
	//rpcAnswersReceived int
	//rpcSentErrors      int
	//rpcSentAnswers     int
	//rpcSentQueries     int
	//tlInAllocated      int
	//tlOutAllocated     int
	//rpcQps             float64
	//defaultRpcFlags    int

	// <<<<<<tl_parse<<<<<<	end
	// >>>>>>crypto_aes>>>>>>	start

	//allocatedAesCrypto     int
	//allocatedAesCryptoTemp int
	//aesPwdHash             string

	//<<<<<<crypto_aes<<<<<<	end

	//>>>>>>crypto_dh>>>>>>	start

	//totDhRounds string

	//<<<<<<crypto_dh<<<<<<	end
	//>>>>>>jobs>>>>>>	start

	//threadAverageIdlePercent string
	//threadRecentIdlePercent  string
	//totThreads               string
	//threadLoadAverageUser    string
	//threadLoadAverageSys     string
	//threadLoadAverage        string
	//threadLoadRecentUser     string
	//threadLoadRecentSys      string
	//threadLoadRecent         string
	//loadAverageUser          float64
	//loadAverageSys           float64
	//loadAverageTotal         float64
	//loadRecentUser           float64
	//loadRecentSys            float64
	//loadRecentTotal          float64
	//maxAverageUser           float64
	//maxAverageSys            float64
	//maxAverageTotal          float64
	//maxRecentUser            float64
	//maxRecentSys             float64
	//maxRecentTotal           float64
	//jobTimersAllocated       int
	//jobsCreated              int
	//jobsActive               int
	//jobsRunning              string
	//jobsAllocatedMemory      int
	//timerOps                 int
	//timerOpsScheduler        int

	//<<<<<<jobs<<<<<<	end
	//>>>>>>mp_queue>>>>>>	start

	//mpqBlocksAllocated         int
	//mpqBlocksAllocatedMax      int
	//mpqBlocksAllocations       int
	//mpqBlocksTrueAllocations   int
	//mpqBlocksWasted            int
	//mpqBlocksPrepared          int
	//mpqSmallBlocksAllocated    int
	//mpqSmallBlocksAllocatedMax int
	//mpqActive                  int
	//mpqAllocated               int

	//<<<<<<mp_queue<<<<<<	end

	//>>>>>>timers>>>>>>	start

	//eventTimerInsertOps int
	//eventTimerRemoveOps int
	//eventTimerAlarms    int
	//totalTimers         int

	//<<<<<<timers<<<<<<	end

	//>>>>>>rpc_targets>>>>>>	start

	//totalRpcTargets              int
	//totalConnectionsInRpcTargets int

	//<<<<<<rpc_targets<<<<<<	end

	//statsGenerateTime                    float64
	//vmsizeBytes                          int
	//vmrssBytes                           int
	//vmdataBytes                          int
	//configFilename                       string
	//configLoadedAt                       int
	//configSize                           int
	//configMd5                            string
	//configAuthClusters                   int
	//workers                              int
	//queriesGet                           int
	//qpsGet                               float64
	//totForwardedQueries                  int
	//expiredForwardedQueries              int
	//droppedQueries                       int
	//totForwardedResponses                int
	//droppedResponses                     int
	//totForwardedSimpleAcks               int
	//droppedSimpleAcks                    int
	//activeRpcsCreated                    int
	//activeRpcs                           int
	//rpcDroppedAnswers                    int
	//rpcDroppedRunning                    int
	//windowClamp                          int
	//totalReadyTargets                    int
	//totalAllocatedTargets                int
	//totalDeclaredTargets                 int
	//totalInactiveTargets                 int
	//TotalConnections int
	//totalEncryptedConnections            int
	//totalAllocatedConnections            int
	//totalAllocatedOutboundConnections    int
	//totalAllocatedInboundConnections     int
	//totalAllocatedSocketConnections      int
	//totalDhConnections                   int
	//totalDhRounds                        string
	TotalSpecialConnections    int
	TotalMaxSpecialConnections int
	//totalAcceptConnectionsFailed         string
	//extConnections                       int
	//extConnectionsCreated                int
	//totalActiveNetworkEvents             int
	//totalNetworkBuffersUsedSize          int
	//totalNetworkBuffersAllocatedBytes    int
	//totalNetworkBuffersUsed              int
	//totalNetworkBufferChunksAllocated    int
	//totalNetworkBufferChunksAllocatedMax int
	//mtprotoProxyErrors                   int
	//connectionsFailedLru                 int
	//connectionsFailedFlood               int
	//httpConnections                      int
	//pendingHttpQueries                   int
	//httpQueries                          int
	//httpBadHeaders                       int
	//httpQps                              float64
	//proxyMode                            int
	//proxyTagSet                          int
	//version                              string
}

func (stats *Stats) UnmarshalText(text []byte) error {
	s := string(text)
	lines := strings.Split(s, "\n")
	for _, v := range lines {
		pair := strings.Split(v, "\t")
		if len(pair) < 2 {
			continue
		}
		key := pair[0]
		value := pair[1]

		f := reflect.ValueOf(stats).Elem().FieldByName(strcase.ToCamel(key))

		if !f.IsValid() || !f.CanSet() {
			continue
		}
		switch f.Kind() {
		default:
			continue
		case reflect.Int:
			i, _ := strconv.ParseInt(value, 10, 64)
			f.SetInt(i)
		case reflect.Float64:
		case reflect.Float32:
			i, _ := strconv.ParseFloat(value, 64)
			f.SetFloat(i)
		}
	}

	return nil
}

func (stats *Stats) ToInfluxFormat() (string) {
	var buffer bytes.Buffer
	buffer.WriteString("proxy")
	buffer.WriteString(" ")

	s := reflect.ValueOf(stats).Elem()

	for i := 0; i < s.NumField(); i += 1 {
		if i != 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(s.Type().Field(i).Name)
		buffer.WriteString("=")
		switch s.Field(i).Kind() {
		default:
			continue
		case reflect.Int:
			buffer.WriteString(fmt.Sprintf("%d", s.Field(i)))
		case reflect.Float32:
		case reflect.Float64:
			buffer.WriteString(fmt.Sprintf("%f", s.Field(i)))

		}
	}

	return buffer.String()

}
