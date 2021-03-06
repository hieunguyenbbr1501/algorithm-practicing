package main

import (
	"fmt"
	"strings"
	
)

func main() {
	var logs = []string{"Lorem", "ipsum", "dolor", "sit", "amet"};
	logs = mergesort(logs, 0, len(logs)-1);
	print(logs);
}

func print(logs []string) {
	for i:=0; i < len(logs); i++ {
		fmt.Print(logs[i] + " ");
	}
	fmt.Println();
}

func mergesort(logs []string, left int, right int) []string {
	print(logs);
	if right - left == 0 {
		return logs;
	}
	
	middle := (right + left)/2;
	
	logs_left := make([]string, middle - left);
	logs_right := make([]string, right - middle + 1);
	
	for i:=left; i < middle; i++ {
		logs_left[i-left] = logs[i-left];
	}
	for i:=middle; i < right; i++ {
		logs_right[i-middle] = logs[i];
	}
	logs_left = mergesort(logs_left, 0, middle - left - 1);
	logs_right = mergesort(logs_right, 0, right - middle);
	
	logs = merge(logs_left, logs_right);
	
	return logs;
	
}

func merge(logs1 []string, logs2 []string)  []string {
	left := 0;
	right := 0;
	logs := make([]string, len(logs1) + len(logs2));
	for i := 0; i < len(logs1) + len(logs2); i++ {
		if left == len(logs1) {
			logs[i] = logs2[right];
			right++;
			continue;
		}
		if right == len(logs2) {
			logs[i] = logs1[left];
			left++;
			continue;
		}
		if strings.Compare(logs1[left], logs2[right]) > 0 {
			logs[i] = logs2[right];
			right++;
			continue;
		}
		if strings.Compare(logs2[right], logs1[left]) > 0 {
			logs[i] = logs1[left];
			left++;
			continue;
		}
	}
	return logs;
}
