const std = @import("std");

const data = @embedFile("input.txt");

pub fn p1() !void {
    const allocator = std.heap.page_allocator;

    const file = try std.fs.cwd().openFile("../input/input.txt", .{});
    defer file.close();

    var buffered = std.io.bufferedReader(file.reader());
    var reader = buffered.reader();

    var arr = std.ArrayList(u8).init(allocator);
    defer arr.deinit();

    var sum: u32 = 0;

    while (true) {
        reader.streamUntilDelimiter(arr.writer(), '\n', null) catch |err| switch (err) {
            error.EndOfStream => break,
            else => return err,
        };

        const line = try arr.toOwnedSlice();
        defer allocator.free(line);

        var numbers = std.ArrayList(u32).init(std.heap.page_allocator);
        defer numbers.deinit();

        for (line) |c| {
            if (std.ascii.isDigit(c)) {
                _ = try numbers.append(@as(u32, @intCast(c - '0')));
            }
        }

        if (numbers.items.len >= 1) {
            sum += 10 * numbers.items[0] + numbers.items[numbers.items.len - 1];
        }
        arr.clearRetainingCapacity();
    }

    std.debug.print("Sum: {d}\n", .{sum});
}

fn p2() !void {
    const file = try std.fs.cwd().openFile("../input/test.txt", .{});
    defer file.close();

    var buffered = std.io.bufferedReader(file.reader());
    var bufreader = buffered.reader();

    var buffer: [1000]u8 = undefined;
    @memset(buffer[0..], 0);

    _ = try bufreader.readUntilDelimiterOrEof(buffer[0..], '\r');

    std.debug.print("{s}\n", .{buffer});
}

fn p3() !void {
    var it = std.mem.tokenize(u8, data, "\n");

    var sum: u32 = 0;
    while (it.next()) |line| {
        var numbers = std.ArrayList(u32).init(std.heap.page_allocator);
        defer numbers.deinit();

        for (line) |c| {
            if (std.ascii.isDigit(c)) {
                _ = try numbers.append(@as(u32, @intCast(c - '0')));
            }
        }

        if (numbers.items.len >= 1) {
            sum += 10 * numbers.items[0] + numbers.items[numbers.items.len - 1];
        }
    }

    std.debug.print("Sum: {d}\n", .{sum});
}

pub fn main() !void {
    try p3();
}
