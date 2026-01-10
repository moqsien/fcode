if [ -z "$TMUX" ] && [ -z "$GHOSTTY" ]; then
    if tmux has-session 2>/dev/null; then
        exec tmux attach
    else
        exec tmux new-session
    fi
fi

# 自动 tmux 分屏逻辑 (仅针对第一个窗口)
if [[ -n "$TMUX" && -z "$SPLITED" ]]; then
    # 获取当前窗口的索引号
    CURRENT_WINDOW=$(tmux display-message -p '#I')

    # 判断是否为第一个窗口 (通常是 0，如果你设置了 base-index 1 则改为 1)
    if [[ "$CURRENT_WINDOW" == "0" ]]; then
        # 检查当前窗口是否只有一个面板，防止重复触发
        if [ "$(tmux list-panes | wc -l)" -eq 1 ]; then
            # 标记已分屏，防止当前窗口的子 shell 再次触发
            export SPLITED=1
            
            # 执行分屏逻辑
            tmux split-window -h -p 30
            tmux split-window -v -p 50
            tmux select-pane -t 0
        fi
    fi
fi
